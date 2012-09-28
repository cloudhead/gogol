/*
 * crayola
 * crayola.c
 */
#include <GL/gl.h>
#include <GL/glext.h>
#include <GL/glut.h>
#include <stdio.h>

#include "crayola.h"

/* Window width & height */
int w = 640, h = 640;

GLuint fbo, fbo_texture,
	   program, fboUnif, hslaUnif,
	   tonemapUnif;

/* These are the current values
 * used to affect the post-processing
 * stage. */
float hsla[]    = {0, 0, 0, 0};
float tonemap[] = {1, 1};

void adjustExp(float exp, float max)
{
	tonemap[0] = exp;
	tonemap[1] = max;
}

void adjustHSL(float h, float s, float l)
{
	hsla[0] = h;
	hsla[1] = s;
	hsla[2] = l;
}

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

	/* Use our framebuffer object for rendering */
	glBindFramebuffer(GL_FRAMEBUFFER, fbo);
	glClear(GL_COLOR_BUFFER_BIT | GL_DEPTH_BUFFER_BIT);

	glLoadIdentity();
	glTranslatef(0, 0, 0);
	glEnable(GL_BLEND);
	glBlendFunc(GL_SRC_ALPHA, GL_ONE_MINUS_SRC_ALPHA);
	glEnable(GL_TEXTURE_2D);
	glClearColor(0.0, 0.0, 0.0, 1.0);

	/* Switch to user-land - rendering from the user side
	 * will happen inside the framebuffer. */
	goDisplay(now);

	/* Activate the post-processing shader program,
	 * and switch back to the screen as render output.
	 *
	 * This means all rendering will be transformed by
	 * the shader program, and output to the screen. */
	glUseProgram(program);
	glUniform1i(fboUnif, 0);
	glUniform4fv(hslaUnif, 1, hsla);
	glUniform2fv(tonemapUnif, 1, tonemap);
	glBindFramebuffer(GL_FRAMEBUFFER, 0);

	glClear(GL_COLOR_BUFFER_BIT | GL_DEPTH_BUFFER_BIT);
	glClearColor(0.0, 0.0, 0.0, 1.0);

	/* Bind the texture which is associated with the
	 * framebuffer (it contains our final image) and
	 * render it to screen. */
	glBindTexture(GL_TEXTURE_2D, fbo_texture);

	glBegin(GL_QUADS);
	  glTexCoord2f(0,   0);   glVertex2f(0, 0);
	  glTexCoord2f(1.0, 0);   glVertex2f(w, 0);
	  glTexCoord2f(1.0, 1.0); glVertex2f(w, h);
	  glTexCoord2f(0,   1.0); glVertex2f(0, h);
	glEnd();

	glUseProgram(0);
	glBindFramebuffer(GL_FRAMEBUFFER, 0);

	glutSwapBuffers();
}

void reshape(int width, int height)
{
	w = width;
	h = height;

	glViewport(0, 0, width, height);
	glMatrixMode(GL_PROJECTION);
	glLoadIdentity();
	glOrtho(0.0, width, height, 0.0, -1.0, 1.0);
	glMatrixMode(GL_MODELVIEW);

	glBindTexture(GL_TEXTURE_2D, fbo_texture);
	glTexImage2D(GL_TEXTURE_2D, 0, GL_RGBA, width, height, 0, GL_RGBA, GL_UNSIGNED_BYTE, NULL);
	glBindTexture(GL_TEXTURE_2D, 0);

	goReshape(width, height);
}

int initProgram()
{
	GLuint vs, fs;
	int validate_ok, link_ok;

	if ((vs = createShader("shader.vert", GL_VERTEX_SHADER))   == 0) return 0;
	if ((fs = createShader("shader.frag", GL_FRAGMENT_SHADER)) == 0) return 0;

	program = glCreateProgram();

	glAttachShader(program, vs);
	glAttachShader(program, fs);

	glLinkProgram(program);
	glGetProgramiv(program, GL_LINK_STATUS, &link_ok);

	if (! link_ok) {
		fprintf(stderr, "glLinkProgram: ");
		printlog(program);
		return 0;
	}
	glValidateProgram(program);
	glGetProgramiv(program, GL_VALIDATE_STATUS, &validate_ok);

	if (! validate_ok) {
		fprintf(stderr, "glValidateProgram: ");
		printlog(program);
	}
	fboUnif     = glGetUniformLocation(program, "fbo_texture");
	hslaUnif    = glGetUniformLocation(program, "hsla_adj");
	tonemapUnif = glGetUniformLocation(program, "tone_adj");

	glUseProgram(program);
	glUniform1i(fboUnif, 0);
	glUseProgram(0);
}

void initFramebuffer()
{
	/* Create framebuffer texture */
	glGenTextures(1, &fbo_texture);
	glBindTexture(GL_TEXTURE_2D, fbo_texture);
	glTexParameteri(GL_TEXTURE_2D, GL_TEXTURE_WRAP_S, GL_CLAMP_TO_EDGE);
	glTexParameteri(GL_TEXTURE_2D, GL_TEXTURE_WRAP_T, GL_CLAMP_TO_EDGE);
	glTexParameteri(GL_TEXTURE_2D, GL_TEXTURE_MAG_FILTER, GL_NEAREST);
	glTexParameteri(GL_TEXTURE_2D, GL_TEXTURE_MIN_FILTER, GL_NEAREST);
	glTexImage2D(GL_TEXTURE_2D, 0, GL_RGBA, w, h, 0, GL_RGBA, GL_UNSIGNED_BYTE, NULL);
	glBindTexture(GL_TEXTURE_2D, 0);

	/* Create framebuffer, associate the above texture with it */
	glGenFramebuffers(1, &fbo);
	glBindFramebuffer(GL_FRAMEBUFFER, fbo);
	glFramebufferTexture2D(GL_FRAMEBUFFER, GL_COLOR_ATTACHMENT0, GL_TEXTURE_2D, fbo_texture, 0);

	GLenum status;
	if ((status = glCheckFramebufferStatus(GL_FRAMEBUFFER)) != GL_FRAMEBUFFER_COMPLETE) {
		fprintf(stderr, "glCheckFramebufferStatus: error %p", status);
		return;
	}
	glBindFramebuffer(GL_FRAMEBUFFER, 0);
}

void init(const char *title)
{
	int argc = 0;

	glutInit(&argc, NULL);
	glutInitDisplayMode(GLUT_RGBA | GLUT_DOUBLE | GLUT_DEPTH);
	glutInitWindowPosition(-1, -1);
	glutInitWindowSize(w, h);
	glutCreateWindow(title);
	glutSetCursor(GLUT_CURSOR_INHERIT);
	glutIgnoreKeyRepeat(1);

	initFramebuffer();
	initProgram();

	glutDisplayFunc(display);
	glutReshapeFunc(reshape);
	glutKeyboardFunc(goKeyboard);
	glutKeyboardUpFunc(goKeyboardUp);
	glutSpecialFunc(goSpecial);
	glutSpecialUpFunc(goSpecialUp);
	glutMotionFunc(goMotion);
	glutPassiveMotionFunc(goMotion);
	glutMouseFunc(goMouse);
	glutEntryFunc(goEntry);
	glutVisibilityFunc(visible);

	glClearColor(1.0f, 1.0f, 1.0f, 0.0f);

	goReady();

	glutMainLoop();

	glDeleteTextures(1, &fbo_texture);
	glDeleteFramebuffers(1, &fbo);
	glDeleteProgram(program);
}
