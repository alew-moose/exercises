#!/usr/bin/env perl
use v5.42;
use utf8;
use List::Util qw(min max);
use DDP;

# my $map = get_input('test-input.txt');
my $map = get_input('input.txt');

# for my $row (@$map) {
#     print '.';
#     print @$row;
#     say '.';
# }

my ($letters, $steps) = travel($map);
say "part 1: $letters";
say "part 2: $steps";

sub travel($map) {
    my ($y, $x) = (0, 0);
    my ($dy, $dx) = (1, 0);
    my @letters;
    my $steps = 0;
    for my $i (0..$#{$map->[0]}) {
        $x = $i, last if $map->[0][$i] eq '|';
    }
    while (true) {
        $y += $dy;
        $x += $dx;
        $steps++;

        # my $r = 3;
        # for my $yc (max($y-$r, 0)..min($y+$r, $#$map)) {
        #     for my $xc (max($x-$r, 0)..min($x+$r, $#{$map->[0]})) {
        #         if ($y == $yc && $x == $xc) {
        #             print 'x';
        #         } else {
        #             print $map->[$yc][$xc];
        #         }
        #     }
        #     say '';
        # }
        # say '';
        # <STDIN>;

        my $t = get_field($map, $y, $x);
        if (!defined($t)) {
            return join('', @letters), $steps;
        }
        if ($t =~ /[A-Z]/) {
            push @letters, $t;
        } elsif ($t eq '+') {
            if ($dy != 0) {
                my $left  = get_field($map, $y, $x-1) // '';
                my $right = get_field($map, $y, $x+1) // '';
                if ($left =~ /[A-Z]/ || $left eq '-') {
                    $dx = -1;
                } elsif ($right =~ /[A-Z]/ || $right eq '-') {
                    $dx = 1;
                } else {
                    die "nowhere to go y=$y x=$x dy=$dy dx=$dx";
                }
                $dy = 0;
            } elsif ($dx != 0) {
                my $up   = get_field($map, $y-1, $x) // '';
                my $down = get_field($map, $y+1, $x) // '';
                if ($up =~ /[A-Z]/ || $up eq '|') {
                    $dy = -1;
                } elsif ($down =~ /[A-Z]/ || $down eq '|') {
                    $dy = 1;
                } else {
                    die "nowhere to go y=$y x=$x dy=$dy dx=$dx";
                }
                $dx = 0;
            } else {
                die 'invalid dy dx';
            }
        } elsif ($t eq ' ') {
            return(join '', @letters), $steps;
        }
    }
}

sub get_field($map, $y, $x) {
    return undef if $y < 0 || $y > $#$map || $x < 0 || $x > $#{$map->[0]};
    return $map->[$y][$x];
}

sub get_input($file) {
    open my $fh, '<', $file or die "$file: $!";
    my @rows;
    while (<$fh>) {
        chomp;
        my @row = split //;
        die 'invalid row' if @rows && $rows[-1]->@* != @row;
        push @rows, \@row;
    }
    close $fh;
    \@rows;
}
