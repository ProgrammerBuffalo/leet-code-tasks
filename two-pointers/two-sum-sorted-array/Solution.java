package az.eldareyvazov;

import java.util.Arrays;

public class Solution {

    public static void main(String[] args) {
        int[] numbers = {2, 3, 6, 8, 10, 12, 16, 17};
        int target = 20;

        System.out.println(Arrays.toString(twoSum(numbers, target)));
    }

    public static int[] twoSum(int[] numbers, int target) {
        for (int sum = 0, index1 = 0, index2 = numbers.length - 1; index1 < index2; sum = 0) {
            sum = numbers[index1] + numbers[index2];

            if (sum < target) {
                index1++;
                continue;
            }

            if (sum > target) {
                index2--;
                continue;
            }

            return new int[]{index1, index2};
        }

        return new int[]{-1, -1};
    }

}
