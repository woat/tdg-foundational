import java.util.ArrayList;

public class LongestWordInSub {

	public static String solution(String s, ArrayList<String> d) {
		String longest = "";
		int lastIdx = 0;
		int i = 0;

		for (String word : d) {
			lastIdx = 0;
			i = 0;

			for (char ch : word.toCharArray()) {
				int idx = s.indexOf(ch, lastIdx);
				if (lastIdx >= s.length() || idx == -1) {
					break;
				}
				lastIdx = idx + 1;
				i += 1;
			}

			if (i == word.length() && i > longest.length()) {
				longest = word;
			}
		}

		return longest;
	}

	public static void main(String[] args) {
		String input0 = "abppplee";
		ArrayList<String> input1 = new ArrayList<>();
		input1.add("able");
		input1.add("ale");
		input1.add("apple");
		input1.add("bale");
		input1.add("kangaroo");

		String answer = solution(input0, input1);
		if (answer != "apple") {
			System.out.println(String.format("FAIL: Expected apple, got %s", answer));
		} else {
			System.out.println(String.format("PASS: Got %s", answer));
		}
	}
}
