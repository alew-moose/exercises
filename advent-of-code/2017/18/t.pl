#!/usr/bin/env perl
use v5.42;
use utf8;
use DDP;

# my $instrs = get_input('test-input-1.txt');
# my $instrs = get_input('test-input-2.txt');
my $instrs = get_input('input.txt');

say 'part 1: ', solve_part_1($instrs);
say 'part 2: ', solve_part_2($instrs);


sub solve_part_1($instrs) {
    my $last_played;
    execute_1($instrs, sub($p) { $last_played = $p; false });
    $last_played;
}

use constant {
    RUNNING => 1,
    STOPPED => 2,
    BLOCKED => 3,
};

sub solve_part_2($instrs) {
    my $snd1;
    my $chan0 = make_chan();
    my $chan1 = make_chan();
    my $prog0 = make_prog(0, $chan0, $chan1, sub {});
    my $prog1 = make_prog(1, $chan1, $chan0, sub { $snd1++ });
    while (true) {
        my ($run0, $run1);
        my $state0 = $prog0->();
        while ($state0 == RUNNING) {
            $run0++;
            $state0 = $prog0->();
        }
        my $state1 = $prog1->();
        while ($state1 == RUNNING) {
            $run1++;
            $state1 = $prog1->();
        }
        last unless $run0 || $run1;
    }
    $snd1;
}

sub make_chan {
    my @buf;
    sub {
        return shift @buf unless @_;
        push @buf, @_;
    };
}

sub make_prog($prog_n, $chan_snd, $chan_rcv, $snd_cb) {
    my $regs = {p => $prog_n};
    my $pc = 0;
    my sub make_reg_arg_sub($op) {
        my $sub = eval sprintf 'sub($reg, $arg) { $regs->{$reg} //= 0; $regs->{$reg} %s arg_val($regs, $arg) }', $op;
        die $@ if $@;
        $sub;
    }
    my $ops = {
        set => make_reg_arg_sub('='),
        add => make_reg_arg_sub('+='),
        mul => make_reg_arg_sub('*='),
        mod => make_reg_arg_sub('%='),
        snd => sub($arg) {
            $chan_snd->(arg_val($regs, $arg));
            $snd_cb->();
        },
        rcv => sub($reg) {
            if (defined(my $val = $chan_rcv->())) {
                $regs->{$reg} = $val;
                return true;
            }
            return false;
        },
        jgz => sub($arg1, $arg2) {
            if (arg_val($regs, $arg1) > 0) {
                $pc += arg_val($regs, $arg2);
                return true;
            }
            return false;
        },
    };
    sub {
        return STOPPED if $pc > $#$instrs;
        my ($op_name, @args) = $instrs->[$pc]->@*;
        die "unknown op $op_name" unless $ops->{$op_name};
        my $res = $ops->{$op_name}(@args);
        if ($op_name eq 'jgz') {
            $pc++ unless $res;
        } elsif ($op_name eq 'rcv') {
            if ($res) {
                $pc++;
            } else {
                return BLOCKED;
            }
        } else {
            $pc++;
        }
        return RUNNING;
    };
}

sub arg_val($regs, $arg) {
    $arg =~ /^[a-z]$/ ? $regs->{$arg} // 0 : $arg;
}

sub execute_1($instrs, $rcv_cb) {
    my $last_played;
    my $pc = 0;
    my $regs = {};
    my sub make_reg_arg_sub($op) {
        my $sub = eval sprintf 'sub($reg, $arg) { $regs->{$reg} //= 0; $regs->{$reg} %s arg_val($regs, $arg) }', $op;
        die $@ if $@;
        $sub;
    }
    my $ops = {
        set => make_reg_arg_sub('='),
        add => make_reg_arg_sub('+='),
        mul => make_reg_arg_sub('*='),
        mod => make_reg_arg_sub('%='),
        snd => sub($arg) {
            $last_played = arg_val($regs, $arg)
        },
        rcv => sub($arg) {
            if (arg_val($regs, $arg) != 0) {
                return $rcv_cb->($last_played);
            } else {
                return true;
            }
        },
        jgz => sub($arg1, $arg2) {
            if (arg_val($regs, $arg1) > 0) {
                $pc += arg_val($regs, $arg2);
                return true;
            } else {
                return false;
            }
        },
    };
    while ($pc < @$instrs) {
        my ($op_name, @args) = $instrs->[$pc]->@*;
        die "unknown op $op_name" unless $ops->{$op_name};
        my $res = $ops->{$op_name}(@args);
        $pc++ unless $op_name eq 'jgz' && $res;
        last if $op_name eq 'rcv' && !$res;
    }
}

sub get_input($file) {
    state $op_arg_re = qr/snd|rcv/;
    state $op_reg_arg_re = qr/set|add|mul|mod/;
    state $op_arg_arg_re = qr/jgz/;
    state $reg_re = qr/[a-z]/;
    state $num_re = qr/(?:-?\d+)/;
    state $arg_re = qr/$reg_re|$num_re/;
    state $instr_re = qr/
        ^(?|
            ($op_arg_re)     \s+ ($arg_re) |
            ($op_reg_arg_re) \s+ ($reg_re) \s+ ($arg_re) |
            ($op_arg_arg_re) \s+ ($arg_re) \s+ ($arg_re)
        )$
    /x;

    open my $fh, '<', $file or die "$file: $!";
    my @instrs;
    while (<$fh>) {
        chomp;
        /$instr_re/ or die "failed to parse $_";
        push @instrs, [$1, $2, $3 // ()];
    }
    close $fh;
    \@instrs;
}
