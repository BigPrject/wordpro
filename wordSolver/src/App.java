
import java.util.*;
import java.io.*;



public class App {

    static HashMap<String,Integer> allWords = new HashMap<>();
    static ArrayList<String> matchingWords = new ArrayList<>();
    static char[][] board = new char[4][4];

    public static void main(String[] args) throws Exception {

        
        System.out.println("Hello, World!");
        Scanner sc = new Scanner(System.in);
        System.out.println("What is your String");
        String text = sc.nextLine();
        char[] textList =  text.toCharArray();
        int charStart = 0;
        for(int i = 0; i<board.length; i++){
            for(int j = 0; j<board[0].length; j++){

                board[i][j] = text.charAt(j+charStart);


            }
            charStart+=4;
        }

        try (        BufferedReader  br = new BufferedReader(new FileReader("src/words.txt")))
        {
           while(br.ready()){
               String current = br.readLine();

                for(int i = 0; i<textList.length; i++){

                    if(current.indexOf(textList[i]) != -1 && current.length() > 3){
                        allWords.put(current, 0);
                    }
                }

               
               
           }
       } catch (Exception e) {
        System.out.println("there's been an error");
           System.out.println(e.getMessage());
       }
       boolean[][] flags = new boolean[4][4];


       for(int i = 0; i<board.length; i++){
        for(int j = 0; j<board[0].length; j++){

            boolean d = dfs(board, i, j, flags, "");


        }

    }
       for (String s : matchingWords) {
        System.out.println(s);
       }


    }

    public static boolean dfs(char[][] board, int row, int column, boolean[][] visted, String word){
        
        if(row < 0 || row >= 4 || column < 0 || column >= 4){
            return false;
        }
        if (visted[row][column]){
            return false;
        }

        char letter = board[row][column];
        word += letter;
        visted[row][column] = true;

        if (checkWord(word)){
            matchingWords.add(word);
        }

        boolean found =  dfs(board, row+1, column, visted, word) ||  dfs(board, row, column+1, visted, word) || dfs(board, row -1, column, visted, word) || dfs(board, row, column -1 , visted, word) || dfs(board, row + 1, column + 1, visted, word) || dfs(board, row -1 , column +1, visted, word) || dfs(board, row + 1, column - 1, visted, word) || dfs(board, row -1, column -1, visted, word);

        visted[row][column] = false;
        return found;

    }


    public static boolean checkWord(String word){
        if(allWords.containsKey(word)){
            return true;
        }
        return false;
    }
}
