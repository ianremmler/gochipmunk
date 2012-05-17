package chipmunk

/*
Copyright (c) 2012 Serge Zirukin

Permission is hereby granted, free of charge, to any person obtaining
a copy of this software and associated documentation files (the
"Software"), to deal in the Software without restriction, including
without limitation the rights to use, copy, modify, merge, publish,
distribute, sublicense, and/or sell copies of the Software, and to
permit persons to whom the Software is furnished to do so, subject to
the following conditions:

The above copyright notice and this permission notice shall be
included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE
LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION
OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION
WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
*/

/*
#include <chipmunk.h>

extern void pointQuery(cpShape *s, void *p);

static void space_point_query(cpSpace *s, cpVect point, cpLayers layers, cpGroup group, void *p) {
  cpSpacePointQuery(s, point, layers, group, pointQuery, p);
}

extern void nearestPointQuery(cpShape *s, cpFloat distance, cpVect point, void *p);

static void space_nearest_point_query(cpSpace *space, cpVect point, cpFloat maxDistance,
                                      cpLayers layers, cpGroup group, void *f) {
  cpSpaceNearestPointQuery(space, point, maxDistance, layers, group, nearestPointQuery, f);
}

extern void segmentQuery(cpShape *s, cpFloat t, cpVect n, void *p);

static void space_segment_query(cpSpace *space,
                                cpVect   start,
                                cpVect   end,
                                cpLayers layers,
                                cpGroup  group,
                                void    *f) {
  cpSpaceSegmentQuery(space, start, end, layers, group, segmentQuery, f);
}

extern void bbQuery(cpShape *s, void *p);

static void space_bb_query(cpSpace *space, cpBB bb, cpLayers layers, cpGroup group, void *f) {
  cpSpaceBBQuery(space, bb, layers, group, bbQuery, f);
}

extern void eachShape_space(cpShape *s, void *p);

static void space_each_shape(cpSpace *space, void *f) {
  cpSpaceEachShape(space, eachShape_space, f);
}

extern void eachConstraint_space(cpConstraint *c, void *p);

static void space_each_constraint(cpSpace *space, void *f) {
  cpSpaceEachConstraint(space, eachConstraint_space, f);
}

extern void eachBody_space(cpBody *b, void *p);

static void space_each_body(cpSpace *space, void *f) {
  cpSpaceEachBody(space, eachBody_space, f);
}
*/
import "C"

import (
  "unsafe"
)

// Space is a basic unit of simulation in Chipmunk.
type spaceBase struct {
  s *C.cpSpace
}

type Space interface {
  Free()
  c() *C.cpSpace
  Iterations() int
  Gravity() Vect
  Damping() float64
  IdleSpeedThreshold() float64
  SleepTimeThreshold() float64
  CollisionSlop() float64
  CollisionBias() float64
  CollisionPersistence() Timestamp
  EnableContactGraph() bool
  UserData() interface{}
  StaticBody() Body
  CurrentTimeStep() float64
  SetIterations(i int)
  SetGravity(Vect)
  SetDamping(float64)
  SetIdleSpeedThreshold(float64)
  SetSleepTimeThreshold(float64)
  SetCollisionSlop(float64)
  SetCollisionBias(float64)
  SetCollisionPersistence(Timestamp)
  SetEnableContactGraph(bool)
  SetUserData(interface{})
  IsLocked() bool
  RemoveCollisionHandler(CollisionType, CollisionType)
  Step(float64)
  UseSpatialHash(float64, int)
  ReindexStatic()
  ReindexShape(Shape)
  ReindexShapesForBody(Body)
  AddShape(Shape) Shape
  AddStaticShape(Shape) Shape
  AddBody(Body) Body
  AddConstraint(Constraint) Constraint
  RemoveShape(Shape)
  RemoveStaticShape(Shape)
  RemoveBody(Body)
  RemoveConstraint(Constraint)
  Contains(ContainedInSpace) bool
  PointQuery(Vect, Layers, Group, PointQuery)
  PointQueryFirst(Vect, Layers, Group) Shape
  NearestPointQuery(Vect, float64, Layers, Group, NearestPointQuery)
  SegmentQuery(Vect, Vect, Layers, Group, SegmentQuery)
  BBQuery(BB, Layers, Group, BBQuery)
  ActivateShapesTouchingShape(Shape)
  EachShape(func(Shape))
  EachConstraint(func(Constraint))
  EachBody(func(Body))
}

type ContainedInSpace interface {
  ContainedInSpace(s Space) bool
}

// NewSpace creates a new space.
func NewSpace() Space {
  return spaceBase{C.cpSpaceNew()}
}

// Free removes a space.
func (s spaceBase) Free() {
  C.cpSpaceFree(s.s)
}

func (s spaceBase) c() *C.cpSpace {
  return s.s
}

func cpSpace(s *C.cpSpace) Space {
  if s != nil {
    return spaceBase{s}
  }

  return nil
}

/////////////////////////////////////////////////////////////////////////////

// Iterations returns the number of iterations to use in the impulse solver (to solve contacts).
func (s spaceBase) Iterations() int {
  return int(C.cpSpaceGetIterations(s.s))
}

// Gravity returns current gravity used when integrating velocity for rigid bodies.
func (s spaceBase) Gravity() Vect {
  return cpVect(C.cpSpaceGetGravity(s.s))
}

// Damping returns the damping rate expressed as the fraction of velocity bodies retain each second.
func (s spaceBase) Damping() float64 {
  return float64(C.cpSpaceGetDamping(s.s))
}

// IdleSpeedThreshold returns speed threshold for a body to be considered idle.
func (s spaceBase) IdleSpeedThreshold() float64 {
  return float64(C.cpSpaceGetIdleSpeedThreshold(s.s))
}

// SleepTimeThreshold returns the time a groups of bodies must remain idle in order to "fall asleep".
func (s spaceBase) SleepTimeThreshold() float64 {
  return float64(C.cpSpaceGetSleepTimeThreshold(s.s))
}

// CollisionSlop returns amount of encouraged penetration between colliding shapes.
func (s spaceBase) CollisionSlop() float64 {
  return float64(C.cpSpaceGetCollisionSlop(s.s))
}

// CollisionBias returns the speed of how fast overlapping shapes are pushed apart.
func (s spaceBase) CollisionBias() float64 {
  return float64(C.cpSpaceGetCollisionBias(s.s))
}

// CollisionPersistence returns the number of frames that contact information should persist.
func (s spaceBase) CollisionPersistence() Timestamp {
  return Timestamp(C.cpSpaceGetCollisionPersistence(s.s))
}

// EnableContactGraph returns true if rebuild of the contact graph during each step is enabled.
func (s spaceBase) EnableContactGraph() bool {
  return 0 != int(C.cpSpaceGetEnableContactGraph(s.s))
}

// UserData returns user defined data.
func (s spaceBase) UserData() interface{} {
  return cpData(C.cpSpaceGetUserData(s.s))
}

// StaticBody returns a dedicated static body for the space.
// You don't have to use it, but because it's memory is managed automatically with the space
// it's very convenient.
// You can set its user data pointer to something helpful if you want for callbacks.
func (s spaceBase) StaticBody() Body {
  return cpBody(C.cpSpaceGetStaticBody(s.s))
}

// CurrentTimeStep returns the current (if you are in a callback from SpaceStep())
// or most recent (outside of a SpaceStep() call) timestep.
func (s spaceBase) CurrentTimeStep() float64 {
  return float64(C.cpSpaceGetCurrentTimeStep(s.s))
}

/////////////////////////////////////////////////////////////////////////////

// SetIterations sets the number of iterations to use in the impulse solver to solve contacts.
func (s spaceBase) SetIterations(i int) {
  C.cpSpaceSetIterations(s.s, C.int(i))
}

// SetGravity sets the gravity to pass to rigid bodies when integrating velocity.
func (s spaceBase) SetGravity(g Vect) {
  C.cpSpaceSetGravity(s.s, g.c())
}

// SetDamping sets the damping rate expressed as the fraction of velocity bodies retain each second.
// A value of 0.9 would mean that each body's velocity will drop 10% per second.
// The default value is 1.0, meaning no damping is applied.
// Note this damping value is different than those of DampedSpring and DampedRotarySpring.
func (s spaceBase) SetDamping(d float64) {
  C.cpSpaceSetDamping(s.s, C.cpFloat(d))
}

// SetIdleSpeedThreshold sets the speed threshold for a body to be considered idle.
// The default value of 0.0 means to let the space guess a good threshold based on gravity.
func (s spaceBase) SetIdleSpeedThreshold(t float64) {
  C.cpSpaceSetIdleSpeedThreshold(s.s, C.cpFloat(t))
}

// SetSleepTimeThreshold sets the time a group of bodies must remain idle in order to fall asleep.
// Enabling sleeping also implicitly enables the the contact graph.
// The default value of math.Inf(1) disables the sleeping algorithm.
func (s spaceBase) SetSleepTimeThreshold(t float64) {
  C.cpSpaceSetSleepTimeThreshold(s.s, C.cpFloat(t))
}

// SetCollisionSlop sets amount of encouraged penetration between colliding shapes.
// Used to reduce oscillating contacts and keep the collision cache warm.
// Defaults to 0.1. If you have poor simulation quality,
// increase this number as much as possible without allowing visible amounts of overlap.
func (s spaceBase) SetCollisionSlop(sl float64) {
  C.cpSpaceSetCollisionSlop(s.s, C.cpFloat(sl))
}

// SetCollisionBias sets the speed of how fast overlapping shapes are pushed apart.
// Expressed as a fraction of the error remaining after each second.
// Defaults to pow(1.0 - 0.1, 60.0) meaning that Chipmunk fixes 10% of overlap each frame at 60Hz.
func (s spaceBase) SetCollisionBias(b float64) {
  C.cpSpaceSetCollisionBias(s.s, C.cpFloat(b))
}

// SetCollisionPersistence sets the number of frames that contact information should persist.
// Defaults to 3. There is probably never a reason to change this value.
func (s spaceBase) SetCollisionPersistence(p Timestamp) {
  C.cpSpaceSetCollisionPersistence(s.s, C.cpTimestamp(p))
}

// SetEnableContactGraph enables a rebuild of the contact graph during each step.
// Must be enabled to use the EachArbiter() method of Body.
// Disabled by default for a small performance boost.
// Enabled implicitly when the sleeping feature is enabled.
func (s spaceBase) SetEnableContactGraph(cg bool) {
  C.cpSpaceSetEnableContactGraph(s.s, boolToC(cg))
}

// SetUserData sets user definable data pointer.
// Generally this points to your game's controller or game state
// so you can access it when given a Space reference in a callback.
func (s spaceBase) SetUserData(data interface{}) {
  C.cpSpaceSetUserData(s.s, dataToC(data))
}

/////////////////////////////////////////////////////////////////////////////

// IsLocked returns true if objects cannot be added/removed inside a callback.
func (s spaceBase) IsLocked() bool {
  return cpBool(C.cpSpaceIsLocked(s.s))
}

// RemoveCollisionHandler unsets a collision handler.
func (s spaceBase) RemoveCollisionHandler(a CollisionType, b CollisionType) {
  C.cpSpaceRemoveCollisionHandler(s.s, C.cpCollisionType(a), C.cpCollisionType(b))
}

// Step makes the space step forward in time by dt seconds.
func (s spaceBase) Step(dt float64) {
  C.cpSpaceStep(s.s, C.cpFloat(dt))
}

// UseSpatialHash switches the space to use a spatial has as it's spatial index.
func (s spaceBase) UseSpatialHash(dim float64, count int) {
  C.cpSpaceUseSpatialHash(s.s, C.cpFloat(dim), C.int(count))
}

// ReindexStatic updates the collision detection info for the static shape in the space.
func (s spaceBase) ReindexStatic() {
  C.cpSpaceReindexStatic(s.s)
}

// ReindexShape updates the collision detection data for a specific shape in the space.
func (s spaceBase) ReindexShape(sh Shape) {
  C.cpSpaceReindexShape(s.s, sh.c())
}

// ReindexShapesForBody updates the collision detection data for all shapes attached to a body.
func (s spaceBase) ReindexShapesForBody(b Body) {
  C.cpSpaceReindexShapesForBody(s.s, b.c())
}

/////////////////////////////////////////////////////////////////////////////

// AddShape adds a collision shape to the simulation.
// If the shape is attached to a static body, it will be added as a static shape.
func (s spaceBase) AddShape(sh Shape) Shape {
  return cpShape(C.cpSpaceAddShape(s.s, sh.c()))
}

// AddStaticShape explicity adds a shape as a static shape to the simulation.
func (s spaceBase) AddStaticShape(sh Shape) Shape {
  return cpShape(C.cpSpaceAddStaticShape(s.s, sh.c()))
}

// AddBody adds a rigid body to the simulation.
func (s spaceBase) AddBody(b Body) Body {
  return Body{b: C.cpSpaceAddBody(s.s, b.b)}
}

// AddConstraint adds a constraint to the simulation.
func (s spaceBase) AddConstraint(c Constraint) Constraint {
  return cpConstraint(C.cpSpaceAddConstraint(s.s, c.c()))
}

// RemoveShape removes a collision shape from the simulation.
func (s spaceBase) RemoveShape(sh Shape) {
  C.cpSpaceRemoveShape(s.s, sh.c())
}

// RemoveStaticShape removes a collision shape added using AddStaticShape() from the simulation.
func (s spaceBase) RemoveStaticShape(sh Shape) {
  C.cpSpaceRemoveStaticShape(s.s, sh.c())
}

// RemoveBody removes a rigid body from the simulation.
func (s spaceBase) RemoveBody(b Body) {
  C.cpSpaceRemoveBody(s.s, b.b)
}

// RemoveConstraint removes a constraint from the simulation.
func (s spaceBase) RemoveConstraint(c Constraint) {
  C.cpSpaceRemoveConstraint(s.s, c.c())
}

/////////////////////////////////////////////////////////////////////////////

// Contains tests if a collision shape, rigid body or a constraint has been added to the space.
func (s spaceBase) Contains(o ContainedInSpace) bool {
  return o.ContainedInSpace(s)
}

// PointQuery is a callback function type for PointQuery function.
type PointQuery func(s Shape)

//export pointQuery
func pointQuery(s *C.cpShape, p unsafe.Pointer) {
  f := *(*PointQuery)(p)
  f(cpShape(s))
}

// PointQuery queries the space at a point and calls a callback function for each shape found.
func (s spaceBase) PointQuery(point Vect, layers Layers, group Group, f PointQuery) {
  C.space_point_query(s.s, point.c(), layers.c(), group.c(), unsafe.Pointer(&f))
}

// PointQueryFirst queries the space at a point and returns
// the first shape found. Returns nil if no shapes were found.
func (s spaceBase) PointQueryFirst(point Vect, layers Layers, group Group) Shape {
  return cpShape(C.cpSpacePointQueryFirst(s.s, point.c(), layers.c(), group.c()))
}

// NearestPointQuery is a callback function type for NearestPointQuery function.
type NearestPointQuery func(s Shape, distance float64, point Vect)

//export nearestPointQuery
func nearestPointQuery(s *C.cpShape, distance C.cpFloat, point C.cpVect, p unsafe.Pointer) {
  f := *(*NearestPointQuery)(p)
  f(cpShape(s), float64(distance), cpVect(point))
}

// NearestPointQuery queries the space at a point and calls a callback function for each shape found.
func (s spaceBase) NearestPointQuery(
  point Vect,
  maxDistance float64,
  layers Layers,
  group Group,
  f NearestPointQuery) {
  C.space_nearest_point_query(
    s.s,
    point.c(),
    C.cpFloat(maxDistance),
    layers.c(),
    group.c(),
    unsafe.Pointer(&f))
}

// SegmentQuery is a query callback function type.
type SegmentQuery func(s Shape, t float64, n Vect)

//export segmentQuery
func segmentQuery(s *C.cpShape, t C.cpFloat, n C.cpVect, p unsafe.Pointer) {
  f := *(*SegmentQuery)(p)
  f(cpShape(s), float64(t), cpVect(n))
}

// SegmentQuery performs a directed line segment query (like a raycast)
// against the space calling a callback function for each shape intersected.
func (s spaceBase) SegmentQuery(start, end Vect, layers Layers, group Group, f SegmentQuery) {
  C.space_segment_query(s.s, start.c(), end.c(), layers.c(), group.c(), unsafe.Pointer(&f))
}

// BBQuery is a rectangle query callback function type.
type BBQuery func(s Shape)

//export bbQuery
func bbQuery(s *C.cpShape, p unsafe.Pointer) {
  f := *(*BBQuery)(p)
  f(cpShape(s))
}

// BBQuery performs a fast rectangle query on the space calling a callback
// function for each shape found.
// Only the shape's bounding boxes are checked for overlap, not their full shape.
func (s spaceBase) BBQuery(bb BB, layers Layers, group Group, f BBQuery) {
  C.space_bb_query(s.s, bb.c(), layers.c(), group.c(), unsafe.Pointer(&f))
}

// ActivateShapesTouchingShape activates body (calls Activate()) of any shape
// that overlaps the given shape.
func (s spaceBase) ActivateShapesTouchingShape(sh Shape) {
  C.cpSpaceActivateShapesTouchingShape(s.s, sh.c())
}

/////////////////////////////////////////////////////////////////////////////

//export eachShape_space
func eachShape_space(sh *C.cpShape, p unsafe.Pointer) {
  f := *(*func(Shape))(p)
  f(cpShape(sh))
}

//export eachConstraint_space
func eachConstraint_space(c *C.cpConstraint, p unsafe.Pointer) {
  f := *(*func(Constraint))(p)
  f(cpConstraint(c))
}

//export eachBody_space
func eachBody_space(b *C.cpBody, p unsafe.Pointer) {
  f := *(*func(Body))(p)
  f(cpBody(b))
}

// EachShape calls a callback function on each shape in the space.
func (s spaceBase) EachShape(iter func(Shape)) {
  p := unsafe.Pointer(&iter)
  C.space_each_shape(s.s, p)
}

// EachConstraint calls a callback function on each constraint in the space.
func (s spaceBase) EachConstraint(iter func(Constraint)) {
  p := unsafe.Pointer(&iter)
  C.space_each_constraint(s.s, p)
}

// EachBody calls a callback function on each body in the space.
func (s spaceBase) EachBody(iter func(Body)) {
  p := unsafe.Pointer(&iter)
  C.space_each_body(s.s, p)
}
