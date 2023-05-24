package com.mycompany.projetop2.Classe;
import java.util.Scanner;
import java.util.Random;
import java.io.IOException;

/**
 *
 * @author Pedro Henrique
 */

public class Menu {
    
    public int opcao;
    public int numPla;
    public int numEsc;
    public int restoDiv;
    public int randNum;
    public int tentativas;
    public String nome;
    public char novamente;
    Scanner entrada = new Scanner(System.in);
    Random gerador = new Random();
    
public void entradaDados(){
    
    System.out.print("Por favor, entre com o seu nome: ");
    nome = entrada.nextLine();
    System.out.println("");
}

public int menu(){
    do {
        
        System.out.println(this.nome + " , seja bem-vindo ao jogo de adivinhação do Cazé!");
        System.out.println("");
        System.out.println("1. Iniciar o jogo");
        System.out.println("2. Sair do jogo");
        System.out.println("");
        System.out.print("Escolha uma opção: ");
        opcao = entrada.nextInt();
        
        switch (opcao) {
            
            case 1 -> {
               clearConsole();
                
                System.out.print("Escolha um número inteiro qualquer para o ser o número máximo gerado possível: ");
                numEsc = entrada.nextInt();
                
                randNum = gerador.nextInt(numEsc + 1); 
                numPla = -1;
                
                tentativas = 0;
                
               clearConsole();
                
                while (numPla != randNum) {
                    System.out.print("Digite o seu palpite (entre 0 e " + this.numEsc + "): ");
                    numPla = entrada.nextInt();
                    tentativas++;
                    System.out.println("");
                
                
                if (numPla < randNum) {
                    System.out.println("Palpite baixo! Tente novamente!");
                    System.out.println("");
                }
                
                else if (numPla > randNum) {
                    System.out.println("Palpite alto! Tente novamente!");
                    System.out.println("");
                }
                
                else {
                    System.out.println(this.nome + " , parabéns! Você acertou em " + this.tentativas + " tentativas!");
                    System.out.println("");
                }
              }
            }
                
            case 2 -> {
                System.out.println("Finalizando o programa! Obrigado por jogar!");
                return 0;
            } 
            default -> {
                clearConsole();
                System.out.println("Opção inválida!");
            }
                
        }
        System.out.println("Deseja jogar novamente? (S/N): ");
        novamente = entrada.next().charAt(0);
    } while(Character.toUpperCase(novamente) == 'S' || Character.toLowerCase(novamente) == 's');
    
    entrada.close();
    return 0;
}

public static void clearConsole(){
    //Clears Screen in java
    try {
        if (System.getProperty("os.name").contains("Windows"))
            new ProcessBuilder("cmd", "/c", "cls").inheritIO().start().waitFor();
        else
            Runtime.getRuntime().exec("clear");
    } catch (IOException | InterruptedException ex) {}
}
}





