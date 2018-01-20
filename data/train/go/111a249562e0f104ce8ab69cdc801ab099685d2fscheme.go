package scheme

func LoadAll(path string) {
	LoadAchievement(path)
	LoadAward(path)
	LoadBuilding(path)
	LoadBuildingLvUp(path)
	LoadCommon(path)
	LoadDecoration(path)
	LoadDrop(path)
	LoadFb(path)
	LoadGLocalization(path)
	LoadHero(path)
	LoadHeroCreatePlan(path)
	LoadHeroLvUp(path)
	LoadHeroRank(path)
	LoadHeroStageUp(path)
	LoadInitialMap(path)
	LoadItem(path)
	LoadKingLvUp(path)
	LoadKingSkill(path)
	LoadKingSkillCost(path)
	LoadKingSkillLvUp(path)
	LoadMagicHeroCreate(path)
	LoadMission(path)
	LoadNPC(path)
	LoadObstacle(path)
	LoadPvp(path)
	LoadRoleLvUp(path)
	LoadSkill(path)
	LoadSkillFlash(path)
	LoadSkillLvUp(path)
	LoadSkillStageUp(path)
	LoadSoldier(path)
	LoadSoldierLvUp(path)
	LoadSoldierStageUp(path)
	LoadStage(path)
	LoadStageTask(path)
	LoadSign(path)
	LoadMall(path)
}

func Process() {
	BuildingLvUpProcess()
	HeroCreatePlanProcess()
	HeroLvUpProcess()
	HeroStageUpProcess()
	HeroRankProcess()
	DropProcess()
	AwardProcess()
	SkillLvUpProcess()
	SignProcess()
	SodierLvUpProcess()
	SoldierStageUpProcess()
}
