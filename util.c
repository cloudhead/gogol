/*
 * crayola
 * util.c
 */
#include <GL/gl.h>
#include <stdlib.h>
#include <stdio.h>

char *readfile(const char *path)
{
	FILE   *fp;
	char   *buffer;
	size_t  size;

	if (! (fp = fopen(path, "rb"))) {
		fprintf(stderr, "error reading %s\n", path);
		exit(1);
	}

	fseek(fp, 0L, SEEK_END);
	size = ftell(fp);
	rewind(fp);

	buffer = malloc(size + 1);

	fread(buffer, size, 1, fp);
	fclose(fp);

	buffer[size] = '\0';

	return buffer;
}

/*
 * Display compilation errors from the OpenGL shader compiler
 */
void printlog(GLuint object)
{
	GLint len = 0;
	char *log;

	if (glIsShader(object)) {
		glGetShaderiv(object, GL_INFO_LOG_LENGTH, &len);
	} else if (glIsProgram(object)) {
		glGetProgramiv(object, GL_INFO_LOG_LENGTH, &len);
	} else {
		fprintf(stderr, "error: not a shader or a program\n");
		return;
	}
	log = malloc(len);

	if (glIsShader(object))
		glGetShaderInfoLog(object, len, NULL, log);
	else if (glIsProgram(object))
		glGetProgramInfoLog(object, len, NULL, log);

	fprintf(stderr, "%s", log);
	free(log);
}

/*
 * Compile the shader from file `filename`
 */
GLuint createShader(const char* filename, GLenum type)
{
	GLchar *source = readfile(filename);

	if (source == NULL) {
		fprintf(stderr, "error opening %s", filename);
		return 0;
	}
	GLuint res = glCreateShader(type);
	const GLchar* sources[] = { source };

	glShaderSource(res, 1, sources, NULL);

	free(source);

	glCompileShader(res);

	GLint compile_ok = GL_FALSE;

	glGetShaderiv(res, GL_COMPILE_STATUS, &compile_ok);

	if (compile_ok == GL_FALSE) {
		fprintf(stderr, "%s: ", filename);
		printlog(res);
		glDeleteShader(res);
		return 0;
	}
	return res;
}

