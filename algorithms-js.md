Алгоритм кадане
Применяется для нахождения maximum subarray sum
Хорош в низкой сложности BigO(N)
https://math4everyone.info/blog/2020/12/29/poisk-maksimalnoj-summy-posledovatelnyh-elementov-massiva-algoritm-kadane-dinamicheskoe-programmirovanie/

var maxSequence = function(arr){
    var currentSum = 0;
    return arr.reduce(function(maxSum, number){
        currentSum = Math.max(currentSum+number, 0);
        return Math.max(currentSum, maxSum);
    }, 0);
}
