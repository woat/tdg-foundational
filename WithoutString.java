import java.util.List;
import java.util.Arrays;
import java.util.stream.Collectors;

public class WithoutString {
	public static String solution(String base, String remove) {
		String ans = "";

		int i = 0;
		while(i < base.length()) {
			if (base.charAt(i) == remove.charAt(0) && i + remove.length() < base.length()) {
				if (base.substring(i, i + remove.length()).equals(remove)) {
					i += remove.length();
					continue;
				}
				i++;
			}

			ans += base.charAt(i);
			i += 1;
		}

		return ans;
	}

	public static void main(String[] args) {
		String answer = solution("Hello there", "llo");
		if (!answer.equals("He there")) {
			System.out.printf("FAIL: expected He There, got %s\n", answer);
		} else {
			System.out.printf("PASS: got %s\n", answer);
		}
	}
}
