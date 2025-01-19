package org.example;

public class Main {


    public static void main(String[] args) {
        System.err.println(reverse(1534236469));
    }

    public static int reverse(int x) {
        int absX = Math.abs(x);

        int res = 0;
        int digits = 0;

        for (int tempX = absX; tempX > 0; digits++)
            tempX /= 10;

        if (digits >= 10 && x % 10 > 2)
            return 0;

        for (int i = 1, stepRes = 0; digits >= 0; i *= 10, digits--) {
            stepRes = ((absX / i) % 10) * (int)Math.pow(10, digits - 1);

            if (stepRes + res < 0)
                return 0;

            res = stepRes + res;
        }

        return x > 0 ? res : res * (-1);
    }

}