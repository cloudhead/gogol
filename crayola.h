/*
 * crayola
 * crayola.h
 */
void init       (const char *title);
void adjustHSL  (float h, float s, float l);
void adjustExp  (float, float);

/*
 * go callback functions, defined in crayola.go
 */
extern void goDisplay   (int now);
extern void goReshape   (int w, int h);
extern void goKeyboard  (unsigned char key, int x, int y);
extern void goKeyboardUp(unsigned char key, int x, int y);
extern void goSpecial   (int key, int x, int y);
extern void goSpecialUp (int key, int x, int y);
extern void goMouse     (int b, int s, int x, int y);
extern void goMotion    (int x, int y);
extern void goEntry     (int e);
extern void goReady     ();

