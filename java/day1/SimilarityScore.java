package day1;

import java.io.FileInputStream;
import java.io.FileNotFoundException;
import java.io.InputStreamReader;
import java.util.*;

public class SimilarityScore {
    public static void main(String[] args) throws FileNotFoundException {
        FileInputStream in = new FileInputStream("day1/input.txt");
        Scanner scanner = new Scanner(new InputStreamReader(in));
        List<Long> leftNumbers = new ArrayList<>();
        List<Long> rightNumbers = new ArrayList<>();
        Map<Long, Integer> numberOfAppearencesForNumber = new HashMap<>();
        long similarityScore = 0L;
        String line;
        while (scanner.hasNext()) {
            line = scanner.nextLine();
            String[] splitNumbers = line.split("\\s+");
            Long leftNum = Long.parseLong(splitNumbers[0]);
            Long rightNum = Long.parseLong(splitNumbers[1]);
            leftNumbers.add(leftNum);
            rightNumbers.add(rightNum);
        }
        int numApperences;
        for (Long leftNumber : leftNumbers) {
            numApperences = 0;
            if (numberOfAppearencesForNumber.containsKey(leftNumber)) {
                numApperences = numberOfAppearencesForNumber.get(leftNumber);
            } else {
                for (Long rightNumber : rightNumbers) {
                    if (leftNumber.equals(rightNumber)) {
                        ++numApperences;
                    }
                }
                numberOfAppearencesForNumber.put(leftNumber, numApperences);
            }
            similarityScore += (leftNumber * numApperences);
        }
        System.out.format("Result: %d\n", similarityScore);
    }
}
