#include "includes/Gun.h"
#include "includes/Log.h"
#include "includes/Soldier.h"

void SoliderShoot() {
    Solider sanduo("xusanduo");
    sanduo.AddGun(new Gun("AK47"));
    sanduo.AddBulletToGun(20);
    sanduo.fire();
}
int main() {
    LogInfo("start run cmake learn.");
    SoliderShoot();
    return 0;
}