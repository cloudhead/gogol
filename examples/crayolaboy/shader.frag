#version 130

uniform sampler2D fbo_texture;
varying vec2      f_texcoord;

uniform vec4 hsla_adj;
uniform vec2 tone_adj;

vec4 rgba2hsla(vec4 rgba);
vec4 hsla2rgba(vec4 hsla);

void main(void) {
    vec4 fi = vec4(1.0, 1.0, 0.0, 0.3);

    vec4 color = texture(fbo_texture, f_texcoord);
    vec4 hsla  = rgba2hsla(color);
    vec4 hsla2 = clamp(hsla + hsla_adj, 0.0, 1.0);

    float exp    = tone_adj.x;
    float bright = tone_adj.y;

    color = hsla2rgba(hsla2);

    float Y  = dot(vec4(0.3, 0.59, 0.11, 0.0), color);
    float YD = exp * (exp / bright + 1.0) / (exp + 1.0);

    gl_FragColor = color * YD;
}

float hue(float h, float m1, float m2)
{
    h = h < 0 ? h + 1 : (h > 1 ? h - 1 : h);

    if      (h * 6 < 1) return m1 + (m2 - m1) * h * 6;
    else if (h * 2 < 1) return m2;
    else if (h * 3 < 2) return m1 + (m2 - m1) * (2.0/3.0 - h) * 6;
    else                return m1;
}

vec4 hsla2rgba(vec4 hsla)
{
	float h = hsla.x,
		  s = hsla.y,
		  l = hsla.z,
		  a = hsla.w;

	h = mod(h, 1.0);

	float m2 = l <= 0.5 ? l * (s + 1) : l + s - l * s;
	float m1 = l * 2 - m2;

	return vec4(hue(h + 1.0/3.0, m1, m2),
			    hue(h,           m1, m2),
			    hue(h - 1.0/3.0, m1, m2), a);
}

vec4 rgba2hsla(vec4 rgba)
{
	float r = rgba.r,
		  g = rgba.g,
		  b = rgba.b,
		  a = rgba.a;

	float max = max(max(r, g), b),
		  min = min(min(r, g), b);

	float h, s,
		  l = (max + min) / 2.0,
		  d = max - min;

	if (max == min) {
		h = s = 0;
	} else {
		s = l > 0.5 ? d / (2.0 - max - min) : d / (max + min);

		if      (r == max) h = (g - b) / d + (g < b ? 6.0 : 0);
		else if (g == max) h = (b - r) / d + 2.0;
		else               h = (r - g) / d + 4.0;

		h /= 6;
	}
	return vec4(h, s, l, a);
}
