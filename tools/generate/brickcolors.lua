local default
local brickColors = {}
local i = 0
local bad = 0
while true do
	local bc = BrickColor.new(i)
	if bc.Number == i then
		table.insert(brickColors, bc)
		bad = 0
	else
		default = bc.Number
		bad = bad + 1
	end
	if bad >= 2^10 then
		break
	end
	i = i + 1
end
for i = 1, #brickColors do
	local bc = brickColors[i]
	brickColors[i] = {
		Color = {bc.r, bc.g, bc.b},
		Color8 = {(math.modf(bc.r*255)), (math.modf(bc.g*255)), (math.modf(bc.b*255))},
		IndexColor = BrickColor.new(bc.r, bc.g, bc.b).Number,
		IndexColor3 = BrickColor.new(bc.Color).Number,
		IndexName = BrickColor.new(bc.Name).Number,
		IndexNumber = BrickColor.new(bc.Number).Number,
		Name = bc.Name,
		Number = bc.Number,
	}
end
local palette = {}
local i = 0
while true do
	local ok, bc = pcall(BrickColor.palette, i)
	if not ok then
		break
	end
	table.insert(palette, bc.Number)
	i = i + 1
end
print(game.HttpService:JSONEncode({
	Default = default,
	BrickColors = brickColors,
	Palette = palette,
}))
