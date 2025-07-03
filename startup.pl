#!/usr/bin/perl

use strict;
use warnings;
use HTTP::Tiny;
use Time::HiRes qw(sleep);
use POSIX ":sys_wait_h";
use File::Spec;
use IO::Socket::INET;

=pod 
    A finalidade desse script é faciliitar a inicialização dos serviços.
    O script envia a log dos serviços em arquivos txt para facilidade de leitura.
    
    Data: 25/06/2025

=cut 

print "================================================================================================\n";

my %pids; 

## * Configuração dos serviços 
my %services = (
    go => {
        nome => "Backend Go",
        comando => "./startGo.sh",
        porta => 8000,
        log => "log_go.txt",
    },

    python => {
        nome => "Servidor Python",
        comando => "./startPython.sh",
        porta => 5000,
        log => "log_python.txt",
    }, 
);

## * Função para o encerramento de todos os serviços 
sub safe_exit {
    print "\n Encerrando todos os serviços...\n";
    foreach my $srv (keys %pids) {
        my $pid = $pids{$srv};
        if ($pid && kill 0, $pid){
            print "Encerrando $services{$srv}->{nome} {PID: $pid}..\n";
            kill 'TERM', $pid;
        }
    }
    exit 1;
}

## * Função que espera uma 
## * porta HTTP responder com sucesso
sub wait_for_port {
    ## ? Equivalente a um for each
    my ($porta, $nome) = @_;
    print "Agurdando $nome ficar pronto na porta: $porta...\n";

    my $http = HTTP::Tiny->new;
    my $url = "http://localhost:$porta/check";
    
    ## * Tentativas máximas
    my $max_tries = 30;
    my $try = 0;

    ## ? Enquanto o número de tentativas
    ## ? for menor que o número de tentativas máximas
    ## ? verifique se a respota da url foi de sucesso
    while ($try++ < $max_tries) {
        my $res = $http->get($url);
        if ($res->{success}) {
            print "$nome está pronto!\n";
            return 1;
        }
        sleep 1;
    }

    die "Erro: $nome não respondeu na porta $porta após $max_tries segundos.\n";
}

# * Função para iniciar um comando em 
# * segundo plano
sub start_background {
    my ($srv_key) = @_;
    my $info = $services{$srv_key};

    my $pids = fork();
    die "Erro ao realizar fork: $!" unless defined $pids;

    if ($pids == 0) {
        open STDOUT, '>>', $info->{log};
        open STDERR, '>>', $info->{log};
        exec "/bin/bash", "-c", $info->{comando} or die "Erro ao executar $info->{comando}";
    }

    return $pids;
}

foreach my $srv (qw/go python /) {
    my $info = $services{$srv};
    my $pids = start_background($srv);
    $pids{$srv} = $pids;

    eval {
        wait_for_port($info->{porta}, $info->{nome});
    };

    if($@) {
        print "Erro ao iniciar $info->{nome}: $@\n";
        safe_exit();
    }
}

print "======================== Todos os serviços foram iniciados com sucesso! ========================\n";

