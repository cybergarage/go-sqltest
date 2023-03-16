#!/usr/bin/perl
# Copyright (C) 2020 The go-sqltest Authors. All rights reserved.
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

use strict;
use warnings;
use File::Find;

my $test_dir = "./";

if (1 <= @ARGV){
  $test_dir = $ARGV[0];
}

my @test_files = ();

find(
  sub {
    if (-f && /\.qst$/) {
      push(@test_files, $File::Find::name);
    }
  },
, $test_dir);

print<<HEADER;
// Copyright (C) 2020 The go-sqltest Authors. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

var queryScenarioFiles = []string {
HEADER
foreach my $test_file(@test_files){
  printf("%s\n", $test_file);
}
print<<FOTTER;
}
FOTTER