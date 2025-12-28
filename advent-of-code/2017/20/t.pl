#!/usr/bin/env perl
use v5.42;
use utf8;
use List::Util qw(min);
use DDP;

use constant {
    POS => 0,
    VEL => 1,
    ACC => 2,

    X => 0,
    Y => 1,
    Z => 2,

    A => 0,
    B => 1,
    C => 2,
};

# my $particles = get_input('test-input-1.txt');
# my $particles = get_input('test-input-2.txt');
my $particles = get_input('input.txt');

say 'part 1: ', solve_part_1($particles);
say 'part 2: ', solve_part_2($particles);

sub solve_part_1($particles) {
    my $min_acc;
    my $particle;
    for my $i (0..$#$particles) {
        my $p = $particles->[$i];
        my $acc = vec3_manhattan_length($p->[ACC]);
        if (!defined($min_acc) || $acc < $min_acc) {
            $min_acc = $acc;
            $particle = $i;
        }
    }
    $particle;
}

sub solve_part_2($particles) {
    my %collisions;
    for my $i (0..$#$particles-1) {
        my $me1 = motion_equation($particles->[$i]);
        for my $k ($i+1..$#$particles) {
            my $me2 = motion_equation($particles->[$k]);
            my $ct = collision_time($me1, $me2);
            if (defined $ct) {
                $collisions{$ct}{$i}{$k} = 1;
                $collisions{$ct}{$k}{$i} = 1;
            }
        }
    }
    my @collision_times = sort { $a <=> $b } keys %collisions;
    my %collided;
    for my $ct (@collision_times) {
        for my ($p1, $ps) ($collisions{$ct}->%*) {
            next if $collided{$p1};
            for my $p2 (keys %$ps) {
                next if $collided{$p2};
                $collided{$p2} = 1;
            }
            $collided{$p1} = 1;
        }
    }
    @$particles - %collided;
}

sub motion_equation($p) {
    [
        map {
            [ $p->[ACC][$_]/2, $p->[VEL][$_] + $p->[ACC][$_]/2, $p->[POS][$_] ]
        }
        X, Y, Z
    ];
}

sub particle_to_string($p) {
    sprintf(
        'p=<%d,%d,%d> v=<%d,%d,%d> a=<%d,%d,%d>',
        $p->[POS][X], $p->[POS][Y], $p->[POS][Z],
        $p->[VEL][X], $p->[VEL][Y], $p->[VEL][Z],
        $p->[ACC][X], $p->[ACC][Y], $p->[ACC][Z],
    );
}

sub motion_equation_to_string($me) {
    sprintf(
        'x(t) = %s; y(t) = %s; z(t) = %s',
        quadratic_equation_to_string($me->[X]),
        quadratic_equation_to_string($me->[Y]),
        quadratic_equation_to_string($me->[Z]),
    );
}

sub quadratic_equation_to_string($e) {
    sprintf '%g*t^2 + %g*t + %g', $e->[A], $e->[B], $e->[C];
}

sub motion_equation_diff($me1, $me2) {
    [
        map {
            my ($dim1, $dim2) = ($me1->[$_], $me2->[$_]);
            [
                map {
                    $dim1->[$_] - $dim2->[$_]
                }
                A, B, C
            ]
        }
        X, Y, Z
    ];
}

sub collision_time($me1, $me2) {
    my $ed = motion_equation_diff($me1, $me2);

    my $tx = solve_qe($ed->[X]) // return undef;
    my $ty = solve_qe($ed->[Y]) // return undef;
    my $tz = solve_qe($ed->[Z]) // return undef;

    my @ts = grep { $_ != -1 } ($tx, $ty, $tz);
    return 0 unless @ts;

    my $t1 = shift @ts;
    for my $t2 (@ts) {
        return undef if $t1 != $t2;
    }
    return $t1;
}

sub solve_qe($e) {
    if ($e->[A] == 0 && $e->[B] == 0) {
        if ($e->[C] == 0) {
            return -1;
        }
        return undef;
    }
    if ($e->[A] == 0) {
        my $res = -$e->[C] / $e->[B];
        return $res == int($res) ? $res : undef;
    }
    my $d = $e->[B] * $e->[B] - 4 * $e->[A] * $e->[C];
    return undef if $d < 0;
    my $sqrt_d = sqrt($d);
    my @xs =
        grep { $_ > 0 && $_ == int($_) }
        (
            (-$e->[B] + $sqrt_d) / (2 * $e->[A]),
            (-$e->[B] - $sqrt_d) / (2 * $e->[A]),
        );
    return min @xs;
}

sub vec3_manhattan_length($v) {
    abs($v->[X]) + abs($v->[Y]) + abs($v->[Z]);
}

sub get_input($file) {
    state $num_re = qr/(-?\d+)/;
    state $vec3_re = qr/<$num_re,$num_re,$num_re>/;
    state $particle_re = qr/^p=$vec3_re, v=$vec3_re, a=$vec3_re$/;
    open my $fh, '<', $file or die "$file: $!";
    my @particles;
    while (<$fh>) {
        chomp;
        /$particle_re/ or die "failed to parse $_";
        push @particles, [[$1, $2, $3], [$4, $5, $6], [$7, $8, $9]];
    }
    close $fh;
    \@particles;
}
