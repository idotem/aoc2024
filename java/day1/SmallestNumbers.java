package day1;

import java.io.FileInputStream;
import java.io.FileNotFoundException;
import java.io.InputStreamReader;
import java.util.*;

public class SmallestNumbers {

    public static void main(String[] args) throws FileNotFoundException {
        FileInputStream in = new FileInputStream("day1/input.txt");
        Scanner scanner = new Scanner(new InputStreamReader(in));
        List<Long> leftNumbers = new ArrayList<>();
        List<Long> rightNumbers = new ArrayList<>();
        long totalDistance = 0L;
        String line;
        while (scanner.hasNext()){
            line = scanner.nextLine();
            String [] splitNumbers = line.split("\\s+");
            Long leftNum = Long.parseLong(splitNumbers[0]);
            Long rightNum = Long.parseLong(splitNumbers[1]);
            leftNumbers.add(leftNum);
            rightNumbers.add(rightNum);
        }
        leftNumbers = leftNumbers.stream().sorted().toList();
        rightNumbers = rightNumbers.stream().sorted().toList();
        for (int i = 0; i < leftNumbers.size(); i++) {
            totalDistance += Math.abs(leftNumbers.get(i) - rightNumbers.get(i));
        }
        System.out.format("Result: %d\n", totalDistance);
    }
}
