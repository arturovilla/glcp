# What is this ?

glcp is a small little CLI tool i made becuase manually plugging in the GLSL vlaues from [this](http://dev.thi.ng/gradients/) website is really really annoying.


# Usage

`glcp -values="values go here"`

example: `glcp -values="[[0.500 0.500 0.500] [-1.572 0.500 0.500] [1.000 1.000 1.000] [0.000 0.333 0.667]]"`

expected output:
```
λ ~ glcp values="[[0.500 0.500 0.500] [-1.572 0.500 0.500] [1.000 1.000 1.000] [0.000 0.333 0.667]]"

	**********   COLOR PALLETE FUNCTION   **********

vec3 pallete(float t){
	vec3 a = vec3(0.500, 0.500, 0.500);
	vec3 b = vec3(0.500, 0.500, 0.500);
	vec3 c = vec3(1.000, 1.000, 1.000);
	vec3 d = vec3(0.000, 0.333, 0.667);
	return a + b*cos(6.28318*(c+t+d));
}
```

# How to Install (MacOS only)
1. Have [go](https://go.dev/doc/install) installed
2. Clone repo: `git clone https://github.com/arturovilla/glcp.git`
3. Build the program: `go build -o glcp`
4. Move the binary into your bin: `mv glcp /usr/local/bin/`
5. Give the file needed permissions `chmod +x /usr/local/bin/glcp` 
