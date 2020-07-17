/***
 * Maracatronics Robotics
 * Federal University of Pernambuco (UFPE) at Recife
 * http://www.maracatronics.com/
 *
 * This file is part of Armorial project.
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 ***/

#include "behaviour_test.h"

QString Behaviour_Test::name() {
    return "Behaviour_Test";
}

Behaviour_Test::Behaviour_Test() {
}

void Behaviour_Test::configure() {
    usesSkill(_skill_test = new Skill_Test());
};

void Behaviour_Test::run() {
    std::pair<Position, double> aim = getBestAimPosition();
    Position aimPos = aim.first;
    if(aimPos.isUnknown()){
        //std::cout << "Tava unknown, ajustei pro centro" << std::endl;
        aimPos = loc()->ourGoal();
    }
    double aimAngle = GEARSystem::Angle::toDegrees(aim.second);
    //std::cout << MRCConstants::red << "aimAngle = " << MRCConstants::reset << aimAngle << std::endl;

    if(aimAngle <= 7.0 || player()->playerId() == 1){
        bool isParabolic = false;
        if(player()->playerId() == 3){
            aimPos = PlayerBus::ourPlayer(1)->position();
            QList<quint8> removedList;
            removedList.push_back(1);
            removedList.push_back(3);
            isParabolic = loc()->isVectorObstructed(player()->position(), aimPos, removedList, MRCConstants::_robotRadius * 1.5, false);
        }
        else if(player()->playerId() == 1){
            aimPos = PlayerBus::ourPlayer(5)->position();
            QList<quint8> removedList;
            removedList.push_back(1);
            removedList.push_back(5);
            isParabolic = loc()->isVectorObstructed(player()->position(), aimPos, removedList, MRCConstants::_robotRadius * 1.5, false);
        }
        else if(player()->playerId() == 5){
            aimPos = PlayerBus::ourPlayer(3)->position();
            QList<quint8> removedList;
            removedList.push_back(5);
            removedList.push_back(3);
            isParabolic = loc()->isVectorObstructed(player()->position(), aimPos, removedList, MRCConstants::_robotRadius * 1.5, false);
        }
        // set more low kickPower
        if(isParabolic){
            _skill_test->setKickPower(0.75 * sqrt((player()->distanceTo(aimPos) * 9.8) / sin(2 * GEARSystem::Angle::toRadians(65.0))));
            _skill_test->setIsParabolic(isParabolic);
        }else{
            _skill_test->setKickPower(std::min(6.0, std::max(3.0, 2.0 * player()->distanceTo(aimPos))));
            _skill_test->setIsParabolic(isParabolic);
        }
    }else{
        // max kick power
        _skill_test->setIsParabolic(false);
        _skill_test->setKickPower(MRCConstants::_maxKickPower);
    }

    _skill_test->setAim(aimPos);
    _skill_test->shootWhenAligned(true);
}

std::pair<Position, double> Behaviour_Test::getBestAimPosition(){
    // Margin to avoid select the post as aim (or outside it)
    float postMargin = 0.05*loc()->fieldDefenseWidth()/2;

    // Adjust margin
    if(loc()->ourSide().isRight()) {
        postMargin = -postMargin;
    }

    // shift the position of the post to the center of the goal
    Position theirGoalRightPost = loc()->ourGoalRightPost();
    theirGoalRightPost.setPosition(theirGoalRightPost.x(),
                                   theirGoalRightPost.y() + postMargin,
                                   theirGoalRightPost.z());
    Position theirGoalLeftPost = loc()->ourGoalLeftPost();
    theirGoalLeftPost.setPosition(theirGoalLeftPost.x(),
                                  theirGoalLeftPost.y() - postMargin,
                                  theirGoalLeftPost.z());
    Position posTheirGoal = loc()->ourGoal();

    // get obstacles
    QList<Obstacle> obstacles = FreeAngles::getObstacles(loc()->ball());

    // Shift the obstacles
    QList<Obstacle>::iterator obst;

    for(obst = obstacles.begin(); obst != obstacles.end(); obst++) {
        obst->radius() = 1.2 * MRCConstants::_robotRadius;
        // access the robot=
        PlayerAccess *robot = NULL;

        if(player()->opTeamId() == obst->team()) {
            robot = PlayerBus::theirPlayer(obst->id());
        } else {
            robot = PlayerBus::ourPlayer(obst->id());
        }

        Velocity robotVel;
        if(robot != NULL) {
            robotVel = robot->velocity();
        }

        if(robotVel.abs() > 2*0.015) {
            float time = WR::Utils::distance(loc()->ball(), obst->position())/8.0f;
            if(robotVel.abs()*time > 0.2f) {
                time = 0.2f/robotVel.abs();
            }
            obst->position() = WR::Utils::vectorSum(obst->position(), robotVel, time);
        }
    }

    // get free angle with the shifted obstacles
    QList<FreeAngles::Interval> freeAngles = FreeAngles::getFreeAngles(loc()->ball(), theirGoalRightPost, theirGoalLeftPost, obstacles);

    // Get the largest

    float largestAngle=0, largestMid=0;

    if(freeAngles.size()==0) { // Without free angles
        return std::make_pair(Position(false, 0.0, 0.0, 0.0), 0.0);
    } else {
        for(int i=0; i<freeAngles.size(); i++) {
            float angI = freeAngles.at(i).angInitial();
            float angF = freeAngles.at(i).angFinal();
            WR::Utils::angleLimitZeroTwoPi(&angI);
            WR::Utils::angleLimitZeroTwoPi(&angF);
            float dif = angF - angI;
            WR::Utils::angleLimitZeroTwoPi(&dif);

            if(dif>largestAngle) {
                largestAngle = dif;
                largestMid = angF - dif/2;
            }
        }
    }

    // With free angles

    // Triangle
    float x = posTheirGoal.x() - loc()->ball().x();
    float tg = tan(largestMid);
    float y = tg*x;

    // Impact point
    float pos_y = loc()->ball().y() + y;
    Position impactPos(true, posTheirGoal.x(), pos_y, 0.0);

    // Check if impact pos has enough space for the ball
    bool obstructedWay = loc()->isVectorObstructed(loc()->ball(), impactPos, player()->playerId(), MRCConstants::_ballRadius*1.5, false);

    if(obstructedWay) {
        return std::make_pair(Position(false, 0.0, 0.0, 0.0), 0.0);
    }

    return std::make_pair(impactPos, largestAngle);
}
