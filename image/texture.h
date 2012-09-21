/*
 * lourland
 * texture.h
 */
void textureDraw(unsigned int id,
		int  tw,  int  th,   // Width & height of the texture
		int   x,  int   y,   // Position in the image to draw
		int   w,  int   h,   // Width & height of the image to draw
		float sx, float sy); // Screen position to draw at

unsigned int textureGen(int w, int h, unsigned int *data);


