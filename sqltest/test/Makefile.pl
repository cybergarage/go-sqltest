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

%.go : %.pl \$(wildcard *.qst)
	perl \$< > \$@

HEADER

my $data_type_file = "${script_dir}/data/data_type.pict";
open(IN, $data_type_file) or die "Failed to open $data_type_file: $!";

my @data_types;
while(<IN>){
  chomp($_);
  @data_types = split(/\t/, $_, -1);
  last;
}
close(IN);

my $pict_prefix = "smpl_crud";
print "PICT_TESTS = \\\n";
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

print<<FOTTER;
tests: \${PICT_TESTS}

FOTTER

for (my $n = 0; $n < scalar(@data_types); $n++) {
    my $data_type = lc($data_types[$n]);
    my $scenario_name = "${pict_prefix}_${data_type}.qst";
    print "${scenario_name}: ${pict_prefix}.pl ${data_type_file}\n";
    print "\tperl ${pict_prefix}.pl ${data_type} > ${scenario_name}\n";
    print "\tgit add ${scenario_name}\n";
    print "\tgit commit -m \"Update ${scenario_name}\" ${scenario_name}\n\n";
}
