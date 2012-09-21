/*
 * crayola
 * crayola.c
 */
#include <GL/gl.h>
#include <GL/glext.h>
#include <GL/glut.h>

#include "crayola.h"

void visible(int vis)
{
	if (vis == GLUT_VISIBLE)
		glutIdleFunc(glutPostRedisplay);
	else
		glutIdleFunc(NULL);
}

void display(void)
{
	int now = glutGet(GLUT_ELAPSED_TIME);

	glClear(GL_COLOR_BUFFER_BIT | GL_DEPTH_BUFFER_BIT);

	glLoadIdentity();
	glTranslatef(0, 0, 0);

	glEnable(GL_BLEND);
	glBlendFunc(GL_SRC_ALPHA, GL_ONE_MINUS_SRC_ALPHA);

	glEnable(GL_TEXTURE_2D);

	goDisplay(now);
	glutSwapBuffers();
}

void reshape(int width, int height)
{
	glViewport(0, 0, width, height);
	glMatrixMode(GL_PROJECTION);
	glLoadIdentity();
	glOrtho(0.0, width, height, 0.0, -1.0, 1.0);
	glMatrixMode(GL_MODELVIEW);

	goReshape(width, height);
}

void init(const char *title)
{
	int argc = 0;

	glutInit(&argc, NULL);
	glutInitDisplayMode(GLUT_RGBA | GLUT_DOUBLE | GLUT_DEPTH);
	glutInitWindowPosition(-1, -1);
	glutInitWindowSize(640, 640);
	glutCreateWindow(title);
	glutSetCursor(GLUT_CURSOR_INHERIT);
	glutIgnoreKeyRepeat(1);

	glutDisplayFunc(display);
	glutReshapeFunc(reshape);
	glutKeyboardFunc(goKeyboard);
	glutKeyboardUpFunc(goKeyboardUp);
	glutSpecialFunc(goSpecial);
	glutSpecialUpFunc(goSpecialUp);
	glutPassiveMotionFunc(goMotion);
	glutMouseFunc(goMouse);
	glutEntryFunc(goEntry);
	glutVisibilityFunc(visible);

	glClearColor(1.0f, 1.0f, 1.0f, 0.0f);

	goReady();

	glutMainLoop();
}
