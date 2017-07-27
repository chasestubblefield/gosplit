package main

import (
    "fmt"
    "runtime"
    "github.com/go-gl/gl/v4.1-core/gl"
    "github.com/go-gl/glfw/v3.1/glfw"
)

const windowWidth = 320
const windowHeight = 480

func init() {
    runtime.LockOSThread()
}

func doUIStuff() {
    if err := glfw.Init(); err != nil {
        panic(err)
    }
    defer glfw.Terminate()

    glfw.WindowHint(glfw.Resizable, glfw.False)
    // glfw.WindowHint(glfw.Decorated, glfw.False)
    glfw.WindowHint(glfw.ContextVersionMajor, 4)
    glfw.WindowHint(glfw.ContextVersionMinor, 1)
    glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
    glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)
    window, err := glfw.CreateWindow(windowWidth, windowHeight, "GoSplit", nil, nil)
    if err != nil {
        panic(err)
    }
    window.MakeContextCurrent()

    if err := gl.Init(); err != nil {
        panic(err)
    }

    fmt.Println("OpenGL version", gl.GoStr(gl.GetString(gl.VERSION)))
    width, height := window.GetSize()
    fmt.Println("GetSize()", width, height)
    left, top, right, bottom := window.GetFrameSize()
    fmt.Println("GetFrameSize()", left, top, right, bottom)
    width, height = window.GetFramebufferSize()
    fmt.Println("GetFramebufferSize()", width, height)

    for !window.ShouldClose(){
        gl.ClearColor(0, 0, 0, 1)
        gl.Clear(gl.COLOR_BUFFER_BIT)

        window.SwapBuffers()
        glfw.PollEvents()
    }

    fmt.Println("Exiting...")
}
