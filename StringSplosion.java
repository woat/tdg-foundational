public class StringSplosion {
	public static String solution(String s) {
		if (s.length() == 1) {
			return s;
		}
		return solution(s.substring(0, s.length() - 1)) + s;
	}

	public static void main(String[] args) {
		String answer = solution("Code");
		if (answer == "Code") {
			System.out.printf("FAIL: expected CCoCodCode, got %s\n", answer);
		} else {
			System.out.printf("PASS: got %s\n", answer);
		}
	}
}
