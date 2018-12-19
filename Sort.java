public class Sort {
	public static int[] solution(int[] a) {
		if (a.length == 0) {
			return a;
		}
		int size = 0;
		for (int i = 0; i < a.length - 1; i++) {
			if (a[i] != a[i + 1]) {
				size += 1;
			}
		}
		int[] ans = new int[size+1];
		int idx = 0;
		for (int i = 0; i < a.length - 1; i++) {
			if (a[i] != a[i + 1]) {
				ans[idx] = a[i];
				idx += 1;
			}
		}
		ans[idx] = a[a.length - 1];
		return ans;
	}

	public static void main(String[] args) {
		int[] answer = solution(new int[] {1, 2, 2});
		if (answer.length != 2) {
			System.out.printf("FAIL: expected [1, 2], got [%d, %d...] len=%d\n",
				 	answer[0], answer[1], answer.length);
		} else {
			System.out.printf("PASS: got [%d, %d]\n", answer[0], answer[1]);
		}
	}
}
