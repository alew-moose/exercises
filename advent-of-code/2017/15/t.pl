#!/usr/bin/env perl
use v5.42;
use utf8;
use DDP;

# test input
# my $init_a = 65;
# my $init_b = 8921;

# input
my $init_a = 883;
my $init_b = 879;

say 'part 1: ', solve_part_1($init_a, $init_b);
say 'part 2: ', solve_part_2($init_a, $init_b);


sub solve_part_1($init_a, $init_b) {
    my $gen_a = make_generator(16807, $init_a);
    my $gen_b = make_generator(48271, $init_b);
    my $eq_cnt = 0;
    for (1..4e7) {
        $eq_cnt++ if lowest_16_bits($gen_a->()) == lowest_16_bits($gen_b->());
    }
    $eq_cnt;
}

sub solve_part_2($init_a, $init_b) {
    my $gen_a = make_generator(16807, $init_a);
    my $gen_b = make_generator(48271, $init_b);
    my $eq_cnt = 0;
    for (1..5e6) {
        my $a = $gen_a->();
        $a = $gen_a->() until $a % 4 == 0;
        my $b = $gen_b->();
        $b = $gen_b->() until $b % 8 == 0;
        $eq_cnt++ if lowest_16_bits($a) == lowest_16_bits($b);
    }
    $eq_cnt;
}

sub lowest_16_bits($n) {
    $n & 0xFFFF;
}

sub make_generator($factor, $init) {
    my $prev = $init;
    sub {
        my $res = $prev * $factor % 2147483647;
        $prev = $res;
        $res;
    };
}

