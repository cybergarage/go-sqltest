#!/usr/bin/perl
# Copyright (C) 2022 The go-sqltest Authors. All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

use strict;
use warnings;
use FindBin;

my $script_dir = $FindBin::Bin;

print<<HEADER;
# Copyright (C) 2022 The go-sqltest Authors. All rights reserved.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#    http:#www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

SHELL := bash

all: embed.go

embed.go : embed.pl \$(wildcard *.qst)
	perl \$< > \$@

tests: \${CRUD_TESTS} \${ORDER_TESTS} \${LIMIT_TESTS} \${AGGR_TESTS} \${MATH_TESTS} \${YCSB_TESTS}

HEADER

#
# smpl_crud_<type>.qst
#

my $data_type_file = "data/data_type.pict";
open(IN, "${script_dir}/${data_type_file}") or die "Failed to open $data_type_file: $!";

my @data_types;
while(<IN>){
  chomp($_);
  @data_types = split(/\t/, $_, -1);
  last;
}
close(IN);

my $pict_prefix = "smpl_crud";
print "CRUD_TESTS = \\\n";
for (my $n = 0; $n < scalar(@data_types); $n++) {
    my $data_type = lc($data_types[$n]);
    my $scenario_name = "${pict_prefix}_${data_type}.qst";
    print "\t${scenario_name}";
    if ($n < ((@data_types) - 1)) {
        print " \\";
    }
    print "\n";
}
print "\n";

for (my $n = 0; $n < scalar(@data_types); $n++) {
    my $data_type = lc($data_types[$n]);
    my $scenario_name = "${pict_prefix}_${data_type}.qst";
    print "${scenario_name}: ${pict_prefix}.pl ${data_type_file}\n";
    print "\tperl ${pict_prefix}.pl ${data_type} > ${scenario_name}\n";
    system("touch ${script_dir}/${scenario_name}");
}

system("touch ${script_dir}/${pict_prefix}.pl");
system("touch ${script_dir}/${data_type_file}");
print "\n";

#
# select_order_<type>.qst
# select_limit_<type>.qst
# func_aggr_<type>.qst
# func_math_<type>.qst
#

my @test_targets = ("ORDER_TESTS", "LIMIT_TESTS", "AGGR_TESTS", "MATH_TESTS");
my @test_prefixes = ("select_order", "select_limit", "func_aggr", "func_math");
my @test_data_types = ("INT", "FLOAT", "DOUBLE");

for (my $n = 0; $n < @test_targets; $n++) {
    my $test_target = $test_targets[$n];
    print "${test_target} = \\\n";
    my $test_prefix = $test_prefixes[$n];
    for (my $t = 0; $t < @test_data_types; $t++) {
        my $data_type = lc($test_data_types[$t]);
        my $test_scenario = "${test_prefix}_${data_type}.qst";
        print "\t${test_scenario}";
        if ($t < ((@test_data_types) - 1)) {
            print " \\";
        }
        print "\n";
    }
    print "\n";
    for (my $t = 0; $t < @test_data_types; $t++) {
        my $data_type = lc($test_data_types[$t]);
        my $test_scenario = "${test_prefix}_${data_type}.qst";
        print "${test_scenario}: ${test_prefix}.pl\n";
        print "\tperl ${test_prefix}.pl ${data_type} > ${test_scenario}\n";
        system("touch ${script_dir}/${test_scenario}");
        system("git add ${script_dir}/${test_scenario}");
        system("touch ${script_dir}/${test_prefix}.pl");
        system("touch ${script_dir}/${test_prefix}.pl");
    }
    print "\n";
}

print<<FOOTER
YCSB_TESTS = \\
	ycsb_workload.qst

ycsb_workload.qst: ycsb_workload.pl
	perl ycsb_workload.pl > ycsb_workload.qst
FOOTER