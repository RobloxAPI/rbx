package class

import (
	"github.com/robloxapi/rbx/desc"
)

var _Root = desc.NewClass("<<<ROOT>>>")

func init() {
	Instance.Define(
		desc.Alias{Name: "archivable", Target: "Archivable"},
		desc.Alias{Name: "className", Target: "ClassName"},
		desc.Alias{Name: "children", Target: "Children"},
		desc.Alias{Name: "clone", Target: "Clone"},
		desc.Alias{Name: "destroy", Target: "Destroy"},
		desc.Alias{Name: "findFirstChild", Target: "FindFirstChild"},
		desc.Alias{Name: "getChildren", Target: "GetChildren"},
		desc.Alias{Name: "isA", Target: "IsA"},
		desc.Alias{Name: "isDescendantOf", Target: "IsDescendantOf"},
		desc.Alias{Name: "remove", Target: "Remove"},
		desc.Alias{Name: "childAdded", Target: "ChildAdded"},
	)
	BodyAngularVelocity.Define(
		desc.Alias{Name: "angularvelocity", Target: "AngularVelocity"},
		desc.Alias{Name: "maxTorque", Target: "MaxTorque"},
	)
	BodyForce.Define(
		desc.Alias{Name: "force", Target: "Force"},
	)
	BodyGyro.Define(
		desc.Alias{Name: "cframe", Target: "CFrame"},
		desc.Alias{Name: "maxTorque", Target: "MaxTorque"},
	)
	BodyPosition.Define(
		desc.Alias{Name: "maxForce", Target: "MaxForce"},
		desc.Alias{Name: "position", Target: "Position"},
		desc.Alias{Name: "lastForce", Target: "LastForce"},
	)
	BodyThrust.Define(
		desc.Alias{Name: "force", Target: "Force"},
		desc.Alias{Name: "location", Target: "Location"},
	)
	BodyVelocity.Define(
		desc.Alias{Name: "maxForce", Target: "MaxForce"},
		desc.Alias{Name: "velocity", Target: "Velocity"},
		desc.Alias{Name: "lastForce", Target: "LastForce"},
	)
	Camera.Define(
		desc.Alias{Name: "focus", Target: "Focus"},
	)
	ClickDetector.Define(
		desc.Alias{Name: "mouseClick", Target: "MouseClick"},
	)
	Controller.Define(
		desc.Alias{Name: "bindButton", Target: "BindButton"},
		desc.Alias{Name: "getButton", Target: "GetButton"},
	)
	Debris.Define(
		desc.Alias{Name: "addItem", Target: "AddItem"},
	)
	Fire.Define(
		desc.Alias{Name: "size", Target: "Size"},
	)
	Humanoid.Define(
		desc.Alias{Name: "maxHealth", Target: "MaxHealth"},
		desc.Alias{Name: "loadAnimation", Target: "LoadAnimation"},
		desc.Alias{Name: "takeDamage", Target: "TakeDamage"},
	)
	InsertService.Define(
		desc.Alias{Name: "loadAsset", Target: "LoadAsset"},
	)
	JointInstance.Define(
		desc.Alias{Name: "part1", Target: "Part1"},
	)
	Lighting.Define(
		desc.Alias{Name: "getMinutesAfterMidnight", Target: "GetMinutesAfterMidnight"},
		desc.Alias{Name: "setMinutesAfterMidnight", Target: "SetMinutesAfterMidnight"},
	)
	Mouse.Define(
		desc.Alias{Name: "hit", Target: "Hit"},
		desc.Alias{Name: "target", Target: "Target"},
		desc.Alias{Name: "keyDown", Target: "KeyDown"},
	)
	BasePart.Define(
		desc.Alias{Name: "brickColor", Target: "BrickColor"},
		desc.Alias{Name: "breakJoints", Target: "BreakJoints"},
		desc.Alias{Name: "getMass", Target: "GetMass"},
		desc.Alias{Name: "makeJoints", Target: "MakeJoints"},
		desc.Alias{Name: "resize", Target: "Resize"},
		desc.Alias{Name: "touched", Target: "Touched"},
	)
	FormFactorPart.Define(
		desc.Alias{Name: "formFactor", Target: "FormFactor"},
	)
	SkateboardPlatform.Define(
		desc.Alias{Name: "equipped", Target: "Equipped"},
		desc.Alias{Name: "unequipped", Target: "Unequipped"},
	)
	Model.Define(
		desc.Alias{Name: "breakJoints", Target: "BreakJoints"},
		desc.Alias{Name: "makeJoints", Target: "MakeJoints"},
		desc.Alias{Name: "move", Target: "Move"},
		desc.Alias{Name: "moveTo", Target: "MoveTo"},
	)
	WorldRoot.Define(
		desc.Alias{Name: "findPartOnRay", Target: "FindPartOnRay"},
		desc.Alias{Name: "findPartsInRegion3", Target: "FindPartsInRegion3"},
	)
	Player.Define(
		desc.Alias{Name: "userId", Target: "UserId"},
		desc.Alias{Name: "loadBoolean", Target: "LoadBoolean"},
		desc.Alias{Name: "loadInstance", Target: "LoadInstance"},
		desc.Alias{Name: "loadNumber", Target: "LoadNumber"},
		desc.Alias{Name: "loadString", Target: "LoadString"},
		desc.Alias{Name: "saveBoolean", Target: "SaveBoolean"},
		desc.Alias{Name: "saveInstance", Target: "SaveInstance"},
		desc.Alias{Name: "saveNumber", Target: "SaveNumber"},
		desc.Alias{Name: "saveString", Target: "SaveString"},
		desc.Alias{Name: "isFriendsWith", Target: "IsFriendsWith"},
		desc.Alias{Name: "waitForDataReady", Target: "WaitForDataReady"},
	)
	Players.Define(
		desc.Alias{Name: "localPlayer", Target: "LocalPlayer"},
		desc.Alias{Name: "numPlayers", Target: "NumPlayers"},
		desc.Alias{Name: "getPlayerFromCharacter", Target: "GetPlayerFromCharacter"},
		desc.Alias{Name: "getPlayers", Target: "GetPlayers"},
		desc.Alias{Name: "playerFromCharacter", Target: "PlayerFromCharacter"},
		desc.Alias{Name: "players", Target: "Players"},
	)
	RunService.Define(
		desc.Alias{Name: "setThrottleFramerateEnabled", Target: "SetThrottleFramerateEnabled"},
	)
	ServiceProvider.Define(
		desc.Alias{Name: "getService", Target: "GetService"},
		desc.Alias{Name: "service", Target: "Service"},
	)
	DataModel.Define(
		desc.Alias{Name: "lighting", Target: "Lighting"},
		desc.Alias{Name: "workspace", Target: "Workspace"},
	)
	Sound.Define(
		desc.Alias{Name: "isPlaying", Target: "IsPlaying"},
		desc.Alias{Name: "pause", Target: "Pause"},
		desc.Alias{Name: "play", Target: "Play"},
		desc.Alias{Name: "stop", Target: "Stop"},
	)
	TestService.Define(
		desc.Alias{Name: "isFeatureEnabled", Target: "IsFeatureEnabled"},
	)
	BoolValue.Define(
		desc.Alias{Name: "changed", Target: "Changed"},
	)
	BrickColorValue.Define(
		desc.Alias{Name: "changed", Target: "Changed"},
	)
	CFrameValue.Define(
		desc.Alias{Name: "changed", Target: "Changed"},
	)
	Color3Value.Define(
		desc.Alias{Name: "changed", Target: "Changed"},
	)
	DoubleConstrainedValue.Define(
		desc.Alias{Name: "changed", Target: "Changed"},
	)
	IntConstrainedValue.Define(
		desc.Alias{Name: "changed", Target: "Changed"},
	)
	IntValue.Define(
		desc.Alias{Name: "changed", Target: "Changed"},
	)
	NumberValue.Define(
		desc.Alias{Name: "changed", Target: "Changed"},
	)
	ObjectValue.Define(
		desc.Alias{Name: "changed", Target: "Changed"},
	)
	RayValue.Define(
		desc.Alias{Name: "changed", Target: "Changed"},
	)
	StringValue.Define(
		desc.Alias{Name: "changed", Target: "Changed"},
	)
	Vector3Value.Define(
		desc.Alias{Name: "changed", Target: "Changed"},
	)
	VirtualInputManager.Define(
		desc.Alias{Name: "sendRobloxEvent", Target: "SendRobloxEvent"},
	)

	_Root.Init()
}
