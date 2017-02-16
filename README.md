# GoGG! (warning. under development)

## Warning
This engine is still under constant development and no stability is guaranteed

## Intro
GoGG! is a game engine being developed for a planned game editor called GG!Edit.

#### Targeted platforms:
- PC:Windows, MacOS, Linux
- Mobile: Android, iOS
- Other: Web, ARM

Currently the engine is actively being developed for Linux and Android untill a stable alpha version is completed.

## Development philosophy
GoGG! has structure its packages to only be included if the game requires the package. This allows the engine to be used for multiple things and developers don't have to import unnecesary packages. Here are a few (rather ambitious) example scenarios:
- Game simply renderering a 3D object on a PC **(Engine, Render, and OpenGL)**
- 2D interface with two buttons playing a sound on a Android tablet **(Engine, Audio, Input, Touch, Render, UI, and OpenGLES)**
- Background running network game for Android handheld **Engine, Network)**
- First-person shooter mutliplayer game on all platforms **(Engine, Network, Audio, Input, Keyboard, Mouse, Joystick, Render, UI, OpenGL, OpenGLES, WebGL)**
- 2D splitscreen tetris for PC **(Engine, Input, Keyboard, Joystick, Render, OpenGL)**
- Panoramic VR video **(Engine, Input, VR, Render, OpenGL)**

## Build and Run
```
go get gomobile
gomobile init
go get github.com/autovelop/golang-gde
./build.sh
./run.sh
```

## Todo / Feature plan
- [x] Prototype
- [x] Building using Linux and run demo
- [x] Add ecs and properly centralize the shared code between gomobile+opengl_2_es and go+opengl_4_1
- [x] Drawing anything better than triangles
- [x] Add keyboard, mouse, and touch support
- [x] Improving shaders (texture support)
- [x] Improving shaders (font support)
- [x] Create a UI system
- [x] Refactor everything
- [x] Do more thorough testing on Android, MacOS, and Windows platform.
- [x] Add WebSocket support (for multiplayer use)
- [ ] Improving shaders (animations)
- [ ] Refactor everything again
- [ ] Write tests
- [ ] Write documentation
- [ ] Setup examples and screenshots (develop a few games. FUN!)
- [ ] Properly rename/rebrand project
- [ ] Marketing
- [ ] Kickoff GG!Edit project
