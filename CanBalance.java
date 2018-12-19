public class CanBalance {
	public static boolean solution(int[] a) {
		int x = 0;
		int y = a.length - 1;

		int xtotal = 0;
		int ytotal = 0;

		while (x < y) {
			xtotal += a[x];
			ytotal += a[y];

			if (xtotal > ytotal) {
				y -= 1;
				xtotal -= a[x];
				continue;
			}

			if (xtotal < ytotal) {
				x += 1;
				ytotal -= a[y];
				continue;
			}

			x += 1;
			y -= 1;
		}

		if (x == y) {
			return false;
		}

		return xtotal == ytotal && x != y;
	}

	public static void main(String[] args) {
		boolean answer = solution(new int[] {1, 1, 1, 2, 1});
		if (answer) {
			System.out.printf("PASS: got %b\n", answer);
		} else {
			System.out.printf("FAIL: expected true, got %b\n", answer);
		}
	}
}
