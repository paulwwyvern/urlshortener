local weighted_ids = {
   {id = "3BsYdLdigj", weight = 100},
   {id = "pvPgKzcNdC", weight = 20},
   {id = "WOQlCdARtt", weight = 20},
   {id = "8EDM9sGHI7", weight = 20},
   {id = "tNOSyeCSFr", weight = 20},
   {id = "6e1Tq8eJgd", weight = 20},
   {id = "gIb8VTucox", weight = 20},
   {id = "et6TAyR6xm", weight = 10},
   {id = "zbbNGsPWde", weight = 10},
   {id = "wGrn7tix6o", weight = 5},
   {id = "4AMXkcW7mP", weight = 1},
   {id = "WFB3x0Vf5o", weight = 1}
}

local ids = {}
for _, entry in ipairs(weighted_ids) do
   for i = 1, entry.weight do
      table.insert(ids, entry.id)
   end
end

function init(args)
   math.randomseed(tonumber(wrk.thread_id or 0) * 1000)
end

request = function()
   local id = ids[math.random(1, #ids)]
   return wrk.format("GET", "/" .. id)
end