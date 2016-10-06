/*
Given an array of integers, return indices of the two numbers such that they add up to a specific target.

You may assume that each input would have exactly one solution.

Example:
Given nums = [2, 7, 11, 15], target = 9,

Because nums[0] + nums[1] = 2 + 7 = 9,
return [0, 1].
*/

import java.util.HashMap;
import java.util.Map;

public class Solution {

    public int[] twoSum(int[] numbers, int target) {

        int[] res = new int[2];

Map map = new HashMap();

for(int i = 0; i < numbers.length; i++){

map.put(numbers[i], i);

}

for(int i = 0; i < numbers.length; i++){

	int gap = target - numbers[i];
	
	if(map.get(gap)!= null && (int)map.get(gap)!=i){
	
		res[0] = i+1;
		
		res[1] = (int)map.get(gap) + 1;
		
		break;
	
	}

}

return res;

    }

}