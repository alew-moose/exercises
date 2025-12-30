#!/usr/bin/env perl
use v5.42;
use utf8;
use DDP;

my $instrs = get_input('input.txt');

say 'part 1: ', solve_part_1($instrs);
say 'part 2: ', solve_part_2($instrs);

sub solve_part_1($instrs) {
    my $pc = 0;
    my $regs = {};
    my $mul_cnt = 0;
    my sub make_reg_arg_sub($op) {
        my $sub = eval sprintf 'sub($reg, $arg) { $regs->{$reg} //= 0; $regs->{$reg} %s arg_val($regs, $arg) }', $op;
        die $@ if $@;
        $sub;
    }
    my $ops = {
        set => make_reg_arg_sub('='),
        sub => make_reg_arg_sub('-='),
        mul => make_reg_arg_sub('*='),
        jnz => sub($arg1, $arg2) {
            if (arg_val($regs, $arg1) != 0) {
                $pc += arg_val($regs, $arg2);
                return true;
            }
            return false;
        },
    };
    while ($pc < @$instrs) {
        my ($op_name, @args) = $instrs->[$pc]->@*;
        die "unknown op $op_name" unless $ops->{$op_name};
        my $res = $ops->{$op_name}(@args);
        $pc++ unless $op_name eq 'jnz' && $res;
        $mul_cnt++ if $op_name eq 'mul';
    }
    $mul_cnt;
}

sub solve_part_2($instrs) {
    my $pc = 0;
    my $regs = {a => 1};
    my sub make_reg_arg_sub($op) {
        my $sub = eval sprintf 'sub($reg, $arg) { $regs->{$reg} //= 0; $regs->{$reg} %s arg_val($regs, $arg) }', $op;
        die $@ if $@;
        $sub;
    }
    my $ops = {
        set => make_reg_arg_sub('='),
        sub => make_reg_arg_sub('-='),
        mul => make_reg_arg_sub('*='),
        jnz => sub($arg1, $arg2) {
            if (arg_val($regs, $arg1) != 0) {
                $pc += arg_val($regs, $arg2);
                return true;
            }
            return false;
        },
    };
    while ($pc < @$instrs) {
        my ($op_name, @args) = $instrs->[$pc]->@*;
        die "unknown op $op_name" unless $ops->{$op_name};
        my $res = $ops->{$op_name}(@args);
        $pc++ unless $op_name eq 'jnz' && $res;
    }
    $regs->{h};
}


sub arg_val($regs, $arg) {
    $arg =~ /^[a-z]$/ ? $regs->{$arg} // 0 : $arg;
}

sub get_input($file) {
    state $op_reg_arg_re = qr/set|sub|mul/;
    state $op_arg_arg_re = qr/jnz/;
    state $reg_re = qr/[a-z]/;
    state $num_re = qr/(?:-?\d+)/;
    state $arg_re = qr/$reg_re|$num_re/;
    state $instr_re = qr/
        ^(?|
            ($op_reg_arg_re) \s+ ($reg_re) \s+ ($arg_re) |
            ($op_arg_arg_re) \s+ ($arg_re) \s+ ($arg_re)
        )$
    /x;
    my @instrs;
    open my $fh, '<', $file or die "$file: $!";
    while (<$fh>) {
        /$instr_re/ or die "failed to parse $_";
        push @instrs, [$1, $2, $3];
    }
    close $fh;
    \@instrs;
}
