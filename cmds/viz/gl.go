package main

import (
	"fmt"
	"go/types"
	"runtime"
	"unsafe"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
)

const (
	windowWidth  = 960
	windowHeight = 540
)

func main() {
	runtime.LockOSThread()

	if err := glfw.Init(); err != nil {
		panic(fmt.Errorf("could not initialize glfw: %v", err))
	}

	glfw.WindowHint(glfw.ContextVersionMajor, 3)
	glfw.WindowHint(glfw.ContextVersionMinor, 3)
	glfw.WindowHint(glfw.Resizable, glfw.True)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

	win, err := glfw.CreateWindow(800, 600, "Hello world", nil, nil)

	if err != nil {
		panic(fmt.Errorf("could not create opengl renderer: %v", err))
	}

	win.MakeContextCurrent()

	if err := gl.Init(); err != nil {
		panic(err)
	}

	//gl.ClearColor(0, 0.5, 1.0, 1.0)

	var i uint32
	gl.GenVertexArrays(1, &i)
	gl.BindVertexArray(i)

	var g_vertex_buffer_data = [3][3]float32{

		{-1.0, -1.0, 0.0},
		{1.0, -1.0, 0.0},
		{0.0, 1.0, 0.0},
	}

	p_buffer_data := unsafe.Pointer(&g_vertex_buffer_data)

	var vertexbuffer uint32

	gl.GenBuffers(1, &vertexbuffer)
	gl.BindBuffer(gl.ARRAY_BUFFER, vertexbuffer)
	gl.BufferData(gl.ARRAY_BUFFER, len(g_vertex_buffer_data), p_buffer_data, gl.STATIC_DRAW)

	p0 := types.Interface{}

	for !win.ShouldClose() {
		//gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
		gl.EnableVertexAttribArray(0)
		gl.BindBuffer(gl.ARRAY_BUFFER, vertexbuffer)
		gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 0, unsafe.Pointer(&p0))
		gl.DrawArrays(gl.TRIANGLES, 0, 3)
		gl.DisableVertexAttribArray(0)
		win.SwapBuffers()
		glfw.PollEvents()
	}
}
