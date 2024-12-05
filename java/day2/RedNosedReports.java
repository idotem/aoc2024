package day2;

import utils.FileUtils;

import java.io.FileNotFoundException;
import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;
import java.util.Scanner;
import java.util.stream.Collectors;

public class RedNosedReports {
    public static void main(String[] args) throws FileNotFoundException {
        Scanner scanner = FileUtils.getScannerFromFilePath("day2/input.txt");
        String reportLine;
        int safeLevels = 0;
        int allLevels = 0;
        while (scanner.hasNext()) {
            reportLine = scanner.nextLine();
            String[] levels = reportLine.split("\\s+");
            if (checkIfLevelIsSafe(levels)) {
                ++safeLevels;
            }
            ++allLevels;
        }
        System.out.format("Result: %d\n", safeLevels);
        System.out.format("All lvls: %d\n", allLevels);
    }

    private static boolean checkIfLevelIsSafe(String[] levelsArray) {
        List<Integer> levels = Arrays.stream(levelsArray).map(Integer::parseInt).collect(Collectors.toList());
        return checkIfLevelIsSafeWhenIncreasing(levels, false) ||
                checkIfLevelIsSafeWhenDecreasing(levels, false);
    }

    private static boolean checkIfLevelIsSafeWhenIncreasing(List<Integer> levels, boolean checkAgainForbidden) {
        Integer currLevel;
        Integer prevLevel = levels.getFirst();
        for (int i = 1; i < levels.size(); i++) {
            currLevel = levels.get(i);
            int differenceInLevels = Math.abs(currLevel - prevLevel);
            if (prevLevel >= currLevel || differenceInLevels > 3 || differenceInLevels < 1) {
                if (checkAgainForbidden) {
                    return false;
                } else {
                    return checkIfLevelIsSafeWhenIncreasing(removeLevelFromList(i-1, levels), true) ||
                            checkIfLevelIsSafeWhenIncreasing(removeLevelFromList(i, levels), true);
                }
            }
            prevLevel = currLevel;
        }
        return true;
    }

    private static boolean checkIfLevelIsSafeWhenDecreasing(List<Integer> levels, boolean checkAgainForbidden) {
        Integer currLevel;
        Integer prevLevel = levels.getFirst();
        for (int i = 1; i < levels.size(); i++) {
            currLevel = levels.get(i);
            int differenceInLevels = Math.abs(prevLevel - currLevel);
            if (prevLevel <= currLevel || differenceInLevels > 3 || differenceInLevels < 1) {
                if (checkAgainForbidden) {
                    return false;
                } else {
                    return checkIfLevelIsSafeWhenDecreasing(removeLevelFromList(i-1, levels), true) ||
                            checkIfLevelIsSafeWhenDecreasing(removeLevelFromList(i, levels), true);
                }
            }
            prevLevel = currLevel;
        }
        return true;
    }

    private static List<Integer> removeLevelFromList(int index, List<Integer> levels) {
        List<Integer> newList = new ArrayList<>(levels);
        newList.remove(index);
        return newList;
    }
}
