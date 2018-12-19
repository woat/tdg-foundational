import java.util.List;
import java.util.Arrays;
import java.util.stream.Collectors;

public class MaxSpan {
	public static int solution(int[] a) {
		List<Integer> al = Arrays.stream(a).boxed().collect(Collectors.toList());

		int max = 0;
		for (Integer num : al) {
			int idx0 = al.indexOf(num);
			int idx1 = al.lastIndexOf(num);
			int total = idx0 == idx1 ? 1 : idx1 - idx0 + 1;
			max = max < total ? total : max;
		}
		return max;
	}

	public static void main(String[] args) {
		int answer = solution(new int[] {1, 4, 2, 1, 4, 4, 4});
		if (answer != 6) {
			System.out.printf("FAIL: expected 6, got %d\n", answer);
		} else {
			System.out.printf("PASS: got %d\n", answer);
		}
	}
}
