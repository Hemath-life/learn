// <!--
// Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.
// You may assume that each input would have exactly one solution, and you may not use the same element twice.
// You can return the answer in any order. -->

// Problem Understanding:
// You are given an array of integers nums and a target integer target.
// Your goal is to return the indices of the two numbers in the array that add up to the target.
// Each input will have exactly one solution, and you cannot use the same element twice.

/**
 * @param {number[]} nums
 * @param {number} target
 * @return {number[]}
 */
var twoSum = function (nums, target) {
  let seen = new Map(); // Map to store number and its index

  for (let i = 0; i < nums.length; i++) {
    const number = nums[i];
    const tar = target - number;

    if (seen.has(tar)) {
      return [seen.get(tar), i];
    }

    seen.set(number, i);
  }
};

const nums = [2, 7, 11, 15];
const target = 9;

console.log(twoSum(nums, target));
