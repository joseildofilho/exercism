return function(n)
    local sound = ""
    if n % 3 == 0 then
        sound = sound .. "Pling"
    end
    if n % 5 == 0 then
        sound = sound .. "Plang"
    end
    if n % 7 == 0 then
        sound = sound .. "Plong"
    end

    if sound == "" then
        return tostring(n)
    end

    return sound
end
