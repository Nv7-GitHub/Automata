#version 330

// Input vertex attributes (from vertex shader)
in vec2 fragTexCoord;
in vec4 fragColor;

// Input uniform values
uniform sampler2D texture0;

// Output fragment color
out vec4 finalColor;

// NOTE: Add here your custom variables
uniform float time = 0;

float rand (float off) {
    return fract(sin(dot(fragTexCoord.xy+off,
                         vec2(12.9898,78.233)))*
        43758.5453123);
}

float greaterThan(float a, float b) {
    if (a > b) {
        return 1;
    }
    return 0;
}

void main()
{      
    //float randVal = rand(time)*3;
    //vec2 coord = fragTexCoord + vec2(greaterThan(randVal, 0), greaterThan(randVal, 1));
    vec2 size = textureSize(texture0, 0);
    vec2 coord = fragTexCoord + 1/size;
    vec4 color = texture(texture0, coord);
    if (color != vec4(0)) {
        finalColor = vec4(0, time, 1, 1);
    } else {
        finalColor = vec4(0);
    }
}