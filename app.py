import sys

vals = sys.argv[1:]
valsList = vals[0].replace(" ",":").replace("[","").replace("]","")[1:len(vals[0])-1]
VALS = valsList.split(":")
print("vec3 pallete(float t){")
print(f"\tvec3 a = vec3({VALS[0]}, {VALS[1]}, {VALS[2]});")
print(f"\tvec3 b = vec3({VALS[3]}, {VALS[4]}, {VALS[5]});")
print(f"\tvec3 c = vec3({VALS[6]}, {VALS[7]}, {VALS[8]});")
print(f"\tvec3 d = vec3({VALS[9]}, {VALS[10]}, {VALS[11]});")
print(f"\treturn a + b*cos(6.28318*(c+t+d));")
print("}")