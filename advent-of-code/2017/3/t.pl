#!/usr/bin/env perl
use v5.42;
use utf8;
use Carp 'croak';
use List::Util 'sum';
use DDP;


my $target = 361527;
say 'part 1: ', solve_part_1($target);
say 'part 2: ', solve_part_2($target);


sub solve_part_1($target) {
    manhattan_dist(0, 0, find_pos($target));
}

sub solve_part_2($target) {
    croak 'invalid target' if $target < 1;
    my ($y, $x, $prev_square_size, $grid) = (0, 0, 1, {0 => {0 => 1}});
    while (true) {
        $x++;
        my $n = write_neighbors_sum($grid, $y, $x);
        return $n if $n > $target;
        for (0..$prev_square_size-1) {
            $y--;
            my $n = write_neighbors_sum($grid, $y, $x);
            return $n if $n > $target;
        }
        for (0..$prev_square_size) {
            $x--;
            my $n = write_neighbors_sum($grid, $y, $x);
            return $n if $n > $target;
        }
        for (0..$prev_square_size) {
            $y++;
            my $n = write_neighbors_sum($grid, $y, $x);
            return $n if $n > $target;
        }
        for (0..$prev_square_size) {
            $x++;
            my $n = write_neighbors_sum($grid, $y, $x);
            return $n if $n > $target;
        }
        $prev_square_size += 2;
    }
}

sub write_neighbors_sum($grid, $y, $x) {
    return $grid->{$y}{$x} = sum(
        $grid->{$y+1}{$x  } // 0,
        $grid->{$y+1}{$x+1} // 0,
        $grid->{$y  }{$x+1} // 0,
        $grid->{$y-1}{$x+1} // 0,
        $grid->{$y-1}{$x  } // 0,
        $grid->{$y-1}{$x-1} // 0,
        $grid->{$y  }{$x-1} // 0,
        $grid->{$y+1}{$x-1} // 0,
    );
}

sub manhattan_dist($y1, $x1, $y2, $x2) {
    abs($y1 - $y2) + abs($x2 - $x1);
}

sub find_pos($target) {
    croak 'invalid target' if $target < 1;
    my ($y, $x, $n, $prev_square_size) = (0, 0, 1, 1);
    return ($y, $x) if $n == $target;
    while (true) {
        $x++;
        $n++;
        return ($y, $x) if $n == $target;
        for (0..$prev_square_size-1) {
            $y--;
            $n++;
            return ($y, $x) if $n == $target;
        }
        for (0..$prev_square_size) {
            $x--;
            $n++;
            return ($y, $x) if $n == $target;
        }
        for (0..$prev_square_size) {
            $y++;
            $n++;
            return ($y, $x) if $n == $target;
        }
        for (0..$prev_square_size) {
            $x++;
            $n++;
            return ($y, $x) if $n == $target;
        }
        $prev_square_size += 2;
    }
}

