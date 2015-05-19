"""Tests for build_package script."""

from mock import patch
import os
from textwrap import dedent
import unittest

from build_package import (
    build_binary,
    CREATE_LXC_TEMPLATE,
    get_args,
    main,
    parse_dsc,
    setup_local,
    setup_lxc,
    SourceFile,
)
from utils import temp_dir


DSC_CONTENT = dedent("""\
    Format: 3.0 (quilt)
    Source: juju-core
    Binary: juju-core, juju, juju-local, juju-local-kvm
    Architecture: any all
    Version: 1.24-beta1-0ubuntu1~14.04.1~juju1
    Maintainer: Curtis Hovey <curtis.hovey@canonical.com>
    Homepage: http://launchpad.net/juju-core
    Standards-Version: 3.9.5
    Build-Depends: debhelper (>= 7.0.50~), golang-go, lsb-release, python
    Package-List:
     juju deb devel extra
     juju-core deb devel extra
     juju-local deb devel extra
     juju-local-kvm deb devel extra
    Checksums-Sha1:
     1234 9876 juju-core_1.24-beta1.orig.tar.gz
     2345 4321 juju-core_1.24-beta1-0ubuntu1~14.04.1~juju1.debian.tar.gz
    Checksums-Sha256:
     3456 9876 juju-core_1.24-beta1.orig.tar.gz
     4567 4321 juju-core_1.24-beta1-0ubuntu1~14.04.1~juju1.debian.tar.gz
    Files:
     5678 9876 juju-core_1.24-beta1.orig.tar.gz
     6789 4321 juju-core_1.24-beta1-0ubuntu1~14.04.1~juju1.debian.tar.gz
    Testsuite: autopkgtest
    """)


def make_source_files(workspace, dsc_name):
    dsc_path = os.path.join(workspace, dsc_name)
    dsc_file = SourceFile(None, None, 'my.dsc', dsc_path)
    orig_file = SourceFile(
        '3456', '9876', 'juju-core_1.24-beta1.orig.tar.gz',
        '%s/juju-core_1.24-beta1.orig.tar.gz' % workspace)
    deb_file = SourceFile(
        '4567', '4321',
        'juju-core_1.24-beta1-0ubuntu1~14.04.1~juju1.debian.tar.gz',
        '%s/juju-core_1.24-beta1-0ubuntu1~14.04.1~juju1.debian.tar.gz' %
        workspace)
    source_files = [dsc_file, orig_file, deb_file]
    for sf in source_files:
        with open(sf.path, 'w') as f:
            if '.dsc' in sf.name:
                f.write(DSC_CONTENT)
            else:
                f.write(sf.name)
    return source_files


class BuildPackageTestCase(unittest.TestCase):

    def test_get_args_binary(self):
        args = get_args(
            ['prog', 'binary', 'my.dsc', '~/workspace', 'trusty', 'i386'])
        self.assertEqual('binary', args.command)
        self.assertEqual('my.dsc', args.dsc)
        self.assertEqual('~/workspace', args.location)
        self.assertFalse(args.verbose)

    def test_main_binary(self):
        with patch('build_package.build_binary', autospec=True,
                   return_value=0) as bb_mock:
            code = main(
                ['prog', 'binary', 'my.dsc', '~/workspace', 'trusty', 'i386'])
        self.assertEqual(0, code)
        bb_mock.assert_called_with(
            'my.dsc', '~/workspace', 'trusty', 'i386', verbose=False)

    def test_build_binary(self):
        with patch('build_package.parse_dsc', autospec=True,
                   return_value=['orig', 'debian']) as pd_mock:
            with patch('build_package.setup_local', autospec=True,
                       return_value='build_dir') as sl_mock:
                with patch('build_package.setup_lxc', autospec=True,
                           return_value='trusty-i386') as l_mock:
                    with patch('build_package.build_in_lxc',
                               autospec=True) as bl_mock:
                        code = build_binary(
                            'my.dsc', '~/workspace', 'trusty', 'i386',
                            verbose=False)
        self.assertEqual(0, code)
        pd_mock.assert_called_with('my.dsc', verbose=False)
        sl_mock.assert_called_with(
            '~/workspace', 'trusty', 'i386', ['orig', 'debian'], verbose=False)
        l_mock.assert_called_with('trusty', 'i386', 'build_dir', verbose=False)
        bl_mock.assert_called_with('trusty-i386', verbose=False)

    def test_parse_dsc(self):
        with temp_dir() as workspace:
            expected_files = make_source_files(workspace, 'my.dsc')
            dsc_path = os.path.join(workspace, 'my.dsc')
            source_files = parse_dsc(dsc_path, verbose=False)
        self.assertEqual(expected_files, source_files)

    def test_setup_local(self):
        with temp_dir() as workspace:
            source_files = make_source_files(workspace, 'my.dsc')
            build_dir = setup_local(
                workspace, 'trusty', 'i386', source_files, verbose=False)
            self.assertEqual(
                os.path.join(workspace, 'juju-build-trusty-i386'),
                build_dir)
            self.assertTrue(os.path.isdir(build_dir))

    def test_setup_lxc(self):
        with patch('subprocess.check_output') as co_mock:
            container = setup_lxc(
                'trusty', 'i386', '/build-dir', verbose=False)
        self.assertEqual('trusty-i386', container)
        lxc_script = CREATE_LXC_TEMPLATE.format(
            container='trusty-i386', series='trusty', arch='i386',
            build_dir='/build-dir')
        co_mock.assert_called_with(['bash', '-c', lxc_script])
