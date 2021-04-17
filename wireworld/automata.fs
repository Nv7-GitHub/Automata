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
    vec2 size = textureSize(texture0, 0);
    vec4 color = texture(texture0, fragTexCoord);
    if (color == vec4(0)) {
        finalColor = vec4(0);
    } else if (color == conducterColor) {
        // Conducters
        int nbs = 0;
        vec4 col;
        for (int x = -1; x <= 1; x++) {
            for (int y = -1; y <= 1; y++) {
                // Neighbors
                col = texture(texture0, fragTexCoord + (vec2(x, y)/size));
                if (col == headColor) {
                    nbs++;
                }
            }
        }

        // Effects of neighbors
        if (nbs == 1 || nbs == 2) {
            finalColor = headColor;
        } else {
            finalColor = conducterColor;
        }
    } else if (color == headColor) {
        finalColor = tailColor;
    } else if (color == tailColor) {
        finalColor = conducterColor;
    }
}