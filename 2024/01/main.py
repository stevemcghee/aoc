import sys

def read_file(filename):
    with open(filename, 'r') as f:
        lines = f.readlines()
    return lines

def process_lines(lines):
    left_list = []
    right_list = []
    for line in lines:
        left, right = line.strip().split()
        left_list.append(int(left))
        right_list.append(int(right))
    left_list.sort()
    right_list.sort()
    total_distance = 0
    for i in range(len(left_list)):
        total_distance += abs(left_list[i] - right_list[i])
    return total_distance

def process_lines_similarity(lines):
    left_list = []
    right_list = []
    right_dict = {}
    for line in lines:
        left, right = line.strip().split()
        left_list.append(int(left))
        right_list.append(int(right))
        if right in right_dict:
            right_dict[right] += 1
        else:
            right_dict[right] = 1
    left_list.sort()

    total_similarity = 0
    for i in range(len(left_list)):
        try:
            if right_dict[str(left_list[i])] > 0:
                total_similarity += left_list[i] * right_dict[str(left_list[i])]
        except KeyError:
            pass
    return total_similarity

if __name__ == "__main__":
    if len(sys.argv) != 2:
        filename = "2024/01/input"
    else: 
        filename = sys.argv[1]
    lines = read_file(filename)
    result = process_lines(lines)
    print(result)
    result = process_lines_similarity(lines)
    print(result)
