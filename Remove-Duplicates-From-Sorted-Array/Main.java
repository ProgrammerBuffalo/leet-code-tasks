package org.example;

import java.util.Arrays;

public class Main {
    public static void main(String[] args) {
        System.err.println(removeDuplicates(new int[]{1, 1}));

    }

    public static int removeDuplicates(int[] nums) {
        int uniqueNumsCount = 0;
        for (int i = 0, j = 0, index_to_change = 0; i < nums.length; i++, j++) {
            if (i < nums.length - 1 && nums[i] == nums[i + 1]) {
                int currentNum = nums[i];
                while ( j < nums.length && nums[j] == currentNum)
                    j++;
                nums[index_to_change++] = nums[j - 1];
                i = j - 1;
            } else {
                nums[index_to_change++] = nums[i];
            }
            uniqueNumsCount++;
        }
        return uniqueNumsCount;
    }
}