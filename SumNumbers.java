import java.util.List;
import java.util.Arrays;
import java.util.stream.Collectors;

public class SumNumbers {
	public static int solution(String s) {
		int ans = 0;
		String running = "";
		
		for (char ch : s.toCharArray()) {
			if (Character.isDigit(ch)) {
				running += ch;
				continue;
			}

			ans += !running.equals("") ? Integer.parseInt(running) : 0;
			running = "";
		}

		ans += !running.equals("") ? Integer.parseInt(running) : 0;
		return ans;
	}

	public static void main(String[] args) {
		int answer = solution("aa11b33");
		if (answer != 44) {
			System.out.printf("FAIL: expected 44, got %d\n", answer);
		} else {
			System.out.printf("PASS: got %d\n", answer);
		}
	}
}
