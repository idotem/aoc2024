package utils;

import java.io.FileInputStream;
import java.io.FileNotFoundException;
import java.io.InputStreamReader;
import java.util.Scanner;

public class FileUtils {
    public static Scanner getScannerFromFilePath(String filePath) throws FileNotFoundException {
        FileInputStream fileInputStream = new FileInputStream(filePath);
        return new Scanner(new InputStreamReader(fileInputStream));
    }
}
