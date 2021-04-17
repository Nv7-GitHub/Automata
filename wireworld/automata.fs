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
uniform vec4 conducterColor = vec4(0);
uniform vec4 headColor = vec4(0);
uniform vec4 tailColor = vec4(0);

float rand (float off) {
    return fract(sin(dot(fragTexCoord.xy+off,
                         vec2(12.9898,78.233)))*
        43758.5453123);
}

float getRand(float off) {
    float randVal = rand(time+off)*3;
    if (randVal > 2) {
        return 1;
    }
    if (randVal > 1) {
        return -1;
    }
    return 0;
}

void main()
{     
    /*vec4 color = texture(texture0, fragTexCoord);
    if (color == vec4(0)) {
        finalColor = vec4(0);
    } else if (color == conducterColor) {
        // Conducters
        finalColor = vec4(0, 1, 1, 1);
    } else if (color == headColor) {
        finalColor = tailColor;
    } else if (color == tailColor) {
        finalColor = conducterColor;
    } else {
        finalColor = fragColor;
    }*/
    finalColor = texture(texture0, fragTexCoord);
}