public class Interpret {
	public static int solution(int value, String[] commands, int[] args) {
		if (commands.length != args.length) {
			return -1;
		}
		int ans = value;
		for (int i = 0; i < commands.length; i++) {
			String cmd = commands[i];
			switch(cmd) {
				case "*":
					ans *= args[i];
					break;
				case "+":
					ans += args[i];
					break;
				case "-":
					ans -= args[i];
					break;
				default:
					return -1;
			}
		}
		return ans;
	}

	public static void main(String[] args) {
		int answer = solution(1, new String[]{"+", "*", "-"}, new int[]{1, 3, 2});
		if (answer != 4) {
			System.out.printf("FAIL: expected 4, got %d\n", answer);
		} else {
			System.out.printf("PASS: got %d\n", answer);
		}
	}
}
