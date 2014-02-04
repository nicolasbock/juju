// Copyright 2013 Joyent Inc.
// Licensed under the AGPLv3, see LICENCE file for details.

package joyent

import (
	"fmt"
	"strings"
	"strconv"

	"launchpad.net/juju-core/environs/config"
	"launchpad.net/juju-core/instance"

	"launchpad.net/gojoyent/cloudapi"
)

const(
	firewallRuleAll 	= "FROM * TO * ALLOW %s PORT %d"
)

// Helper method to create a firewall rule string for the given port
func createFirewallRuleAll(port instance.Port) string {
	return fmt.Sprintf(firewallRuleAll, port.Protocol, port.Number)
}

// Helper method to check if a firewall rule string already exist
func ruleExists(rules []cloudapi.FirewallRule, rule string) (bool, string) {
	for _, r := range rules {
		if strings.EqualFold(r.Rule, rule) {
			return true, r.Id
		}
	}

	return false, ""
}

// Helper method to get port from the given firewall rules
func getPorts(rules []cloudapi.FirewallRule) []instance.Port {
	ports := make([]instance.Port, len(rules))
	for _, r := range rules {
		rule := r.Rule
		p := strings.ToUpper(rule[strings.Index(rule, "ALLOW") + 6:strings.Index(rule,"PORT")-1])
		n, _ := strconv.Atoi(rule[strings.LastIndex(rule, " ") + 1:])
		port := instance.Port{Protocol:p, Number:n}
		ports = append(ports, port)
	}

	return ports
}

func (env *JoyentEnviron) OpenPorts(ports []instance.Port) error {
	if env.Config().FirewallMode() != config.FwGlobal {
		return fmt.Errorf("invalid firewall mode %q for opening ports on environment", env.Config().FirewallMode())
	}

	fwRules, err := env.compute.cloudapi.ListFirewallRules()
	if err != nil {
		return fmt.Errorf("cannot get firewall rules: %v", err)
	}

	for _, p := range ports {
		rule := createFirewallRuleAll(p)
		if e, id := ruleExists(fwRules, rule); e {
			_, err := env.compute.cloudapi.EnableFirewallRule(id)
			if err != nil {
				return fmt.Errorf("couldn't enable rule %s: %v", rule, err)
			}
		} else {
			_, err := env.compute.cloudapi.CreateFirewallRule(cloudapi.CreateFwRuleOpts{
				Enabled:        true,
				Rule:           rule,
			})
			if err != nil {
				return fmt.Errorf("couldn't create rule %s: %v", rule, err)
			}
		}
	}

	logger.Infof("ports %v opened in environment", ports)

	return nil
}

func (env *JoyentEnviron) ClosePorts(ports []instance.Port) error {
	if env.Config().FirewallMode() != config.FwGlobal {
		return fmt.Errorf("invalid firewall mode %q for opening ports on environment", env.Config().FirewallMode())
	}

	fwRules, err := env.compute.cloudapi.ListFirewallRules()
	if err != nil {
		return fmt.Errorf("cannot get firewall rules: %v", err)
	}

	for _, p := range ports {
		rule := createFirewallRuleAll(p)
		if e, id := ruleExists(fwRules, rule); e {
			_, err := env.compute.cloudapi.DisableFirewallRule(id)
			if err != nil {
				return fmt.Errorf("couldn't disable rule %s: %v", rule, err)
			}
		} else {
			_, err := env.compute.cloudapi.CreateFirewallRule(cloudapi.CreateFwRuleOpts{
				Enabled:        false,
				Rule:           rule,
			})
			if err != nil {
				return fmt.Errorf("couldn't create rule %s: %v", rule, err)
			}
		}
	}

	logger.Infof("ports %v closed in environment", ports)

	return nil
}

func (env *JoyentEnviron) Ports() ([]instance.Port, error) {
	if env.Config().FirewallMode() != config.FwGlobal {
		return nil, fmt.Errorf("invalid firewall mode %q for opening ports on environment", env.Config().FirewallMode())
	}

	fwRules, err := env.compute.cloudapi.ListFirewallRules()
	if err != nil {
		return nil, fmt.Errorf("cannot get firewall rules: %v", err)
	}

	return getPorts(fwRules), nil
}
