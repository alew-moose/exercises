#!/usr/bin/env perl
use v5.42;
use utf8;
use DDP;

# my $step = 3; # test input
my $step = 377;

say 'part 1: ', solve_part_1($step);
say 'part 2: ', solve_part_2($step);

sub solve_part_1($step) {
    my $l = [0];
    my $i = 0;
    for (1..2017) {
        my $val = $l->[$i] + 1;
        $i = ($i + $step) % @$l;
        splice @$l, $i+1, $#$l-$i, $val, @$l[$i+1..$#$l];
        $i++;
    }
    $i = ($i + 1) % @$l;
    $l->[$i];
}

use constant {
    VALUE => 0,
    NEXT  => 1,
};

sub insert($l, $v) {
    my $node = [$v, $l->[NEXT]];
    $l->[NEXT] = $node;
}

# ~22min
sub solve_part_2($step) {
    my $l = [0];
    $l->[NEXT] = $l;
    my $len = 1;
    for my $n (1..50_000_000) {
        my $val = $l->[VALUE] + 1;
        for (1..$step % $len) {
            $l = $l->[NEXT];
        }
        insert($l, $val);
        $len++;
        $l = $l->[NEXT];
    }
    while ($l->[VALUE] != 0) {
        $l = $l->[NEXT];
    }
    $l->[NEXT][VALUE];
}
