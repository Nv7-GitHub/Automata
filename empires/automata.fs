#version 330

// Input vertex attributes (from vertex shader)
in vec2 fragTexCoord;
in vec4 fragColor;

// Input uniform values
uniform sampler2D texture0;

// Output fragment color
out vec4 finalColor;

// NOTE: Add here your custom variables

float rand (float off) {
    return fract(sin(dot(fragTexCoord.xy+off,
                         vec2(12.9898,78.233)))*
        43758.5453123);
}

void main()
{   
    vec4 color = texture(texture0, fragTexCoord);
    if (color != vec4(0)) {
        finalColor = vec4(rand(0), rand(1), rand(2), 1);
        return;
    }
    finalColor = vec4(1);
}