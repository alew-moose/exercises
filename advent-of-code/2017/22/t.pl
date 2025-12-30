#!/usr/bin/env perl
use v5.42;
use utf8;
use DDP;

use constant {
    TURN_L => -1,
    TURN_R => 1,

    DIR_U  => 0,
    DIR_R  => 1,
    DIR_D  => 2,
    DIR_L  => 3,

    NODE_CLEAN    => 0,
    NODE_WEAKENED => 1,
    NODE_INFECTED => 2,
    NODE_FLAGGED  => 3,
};

# my $grid = get_input('test-input.txt');
my $grid = get_input('input.txt');

say 'part 1: ', solve_part_1($grid);
say 'part 2: ', solve_part_2($grid);

sub solve_part_1($grid) {
    $grid = {%$grid};
    my $dir = DIR_U;
    my ($y, $x) = (0, 0);
    my $infected_cnt = 0;
    for (1..10_000) {
        my $key = coord_key($y, $x);
        my $node = $grid->{$key} // NODE_CLEAN;
        $dir = turn($dir, $node == NODE_INFECTED ? TURN_R : TURN_L);
        if ($node == NODE_CLEAN) {
            $grid->{$key} = NODE_INFECTED;
            $infected_cnt++;
        } else {
            $grid->{$key} = NODE_CLEAN;
        }
        ($y, $x) = move($y, $x, $dir);
    }
    $infected_cnt;
}

sub solve_part_2($grid) {
    $grid = {%$grid};
    my $dir = DIR_U;
    my ($y, $x) = (0, 0);
    my $infected_cnt = 0;
    for (1..10_000_000) {
        my $key = coord_key($y, $x);
        my $node = $grid->{$key} // NODE_CLEAN;
        if ($node == NODE_CLEAN) {
            $dir = turn($dir, TURN_L);
            $grid->{$key} = NODE_WEAKENED;
        } elsif ($node == NODE_WEAKENED) {
            # no turn
            $grid->{$key} = NODE_INFECTED;
            $infected_cnt++;
        } elsif ($node == NODE_INFECTED) {
            $dir = turn($dir, TURN_R);
            $grid->{$key} = NODE_FLAGGED;
        } elsif ($node == NODE_FLAGGED) {
            # turn 180
            $dir = turn(turn($dir, TURN_R), TURN_R);
            $grid->{$key} = NODE_CLEAN;
        } else {
            die "unknown node state " . $node;
        }
        ($y, $x) = move($y, $x, $dir);
    }
    $infected_cnt;
}

sub coord_key($y, $x) {
    "$y:$x";
}

sub turn($dir, $turn) {
    ($dir + $turn) % 4;
}

sub move($y, $x, $dir) {
    state $dirs = {
        DIR_U() => [-1, 0],
        DIR_R() => [0, 1],
        DIR_D() => [1, 0],
        DIR_L() => [0, -1],
    };
    my $d = $dirs->{$dir} // die "unknown dir $dir";
    ($y + $d->[0], $x + $d->[1]);
}

sub get_input($file) {
    open my $fh, '<', $file or die "$file: $!";
    my @infected;
    my $y = 0;
    while (<$fh>) {
        chomp;
        my @row = split //;
        while (my ($x, $n) = each @row) {
            if ($n eq '#') {
                push @infected, [$y, $x];
            }
        }
        $y++;
    }
    close $fh;
    my $grid = {};
    my $center = int($y / 2);
    for my $n (@infected) {
        my ($abs_y, $abs_x) = @$n;
        my ($rel_y, $rel_x) = ($abs_y - $center, $abs_x - $center);
        $grid->{coord_key($rel_y, $rel_x)} = NODE_INFECTED;
    }
    $grid;
}
