import java.util.Scanner;

public class main {

	public static void main(String[] args) {
		 
		/*int randomNum = (int)(Math.random()*100);
		
		System.out.println(randomNum);
		
		Scanner i = new Scanner(System.in);
		
		int inputNum = i.nextInt();
		
		System.out.println(inputNum);*/
		
		Scanner input = new Scanner(System.in);
		
		System.out.println("Hi, I'm thinking about a number between 0 and 100. Can you guess it?");
		int randomNum = (int)(Math.random()*100);
		
		System.out.println("Enter a number: ");
		int user_input = input.nextInt();
		
		
		while(user_input != randomNum) {
			if(user_input < randomNum) {
				System.out.println("Too Low!");
			}
			else {
				System.out.println("Too High!");
			}
			
			System.out.println("Enter a number: ");
			user_input = input.nextInt();
		}
		
		System.out.println("Congrats, You got it!!!");
	}

	//I don't want to play anymore
}
