public class MakeBricks {
	public static boolean solution(int s, int b, int v) {
		return s / 5 < b ? v % 5 <= s : !(b == 0 && s < v) && v - b * 5 <= s;
	}

	public static void main(String[] args) {
		boolean ans = solution(3, 2, 10);
		if (!ans) {
			System.out.printf("FAIL: expected true, got %b\n", ans);
		} else {
			System.out.printf("PASS: got %b\n", ans);
		}
	}
}
