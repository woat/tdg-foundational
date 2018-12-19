import java.util.HashMap;
import java.util.Map;

public class Mapshare {
	public static Map<String, String> solution(Map<String, String> map) {
		map.remove("c");
		map.compute("b", (k, v) -> map.containsKey("a") ? v = map.get("a") : v);
		return map;
	}

	public static void main(String[] args) {
		Map<String, String> map = new HashMap<>();
		map.put("a", "aaa");
		map.put("c", "ccc");
		map.put("d", "ddd");
		solution(map);
		if (!map.get("b").equals("aaa") && map.containsKey("c")) {
			System.out.printf("FAIL: expected parsed map\n");
		} else {
			System.out.printf("PASS: was parsed\n");
		}
	}
}
